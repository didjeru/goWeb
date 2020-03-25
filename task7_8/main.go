package main

import (
	"./config"
	"./db"
	"./models"
	"./templates"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var conf = config.LoadConfiguration("./config/config.json")
var data = db.Init(conf.Database)
var tmpl = templates.Init()

func main() {

	f, err := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(err)
	}
	if f != nil {
		defer func() {
			if err := f.Close(); err != nil {
				log.Println(err)
			}
		}()
	}
	log.SetOutput(f)

	http.HandleFunc("/new", newPostHandler)
	http.HandleFunc("/edit", editPostHandler)
	http.HandleFunc("/post", getPostHandler)
	http.HandleFunc("/save", savePostHandler)
	http.HandleFunc("/", indexHandler)

	port := conf.Port
	log.Println("Server started at port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getIDFromQuery(req *http.Request) (int, error) {
	keys, ok := req.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		return -1, errors.New("Url Param 'id' is missing: ")
	}
	return strconv.Atoi(keys[0])
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		errorHandler(res, http.StatusNotFound)
		return
	}

	tpl := tmpl.GetTemplateByName("index")
	if err := tpl.Execute(res, data.GetAllPosts()); err != nil {
		log.Println(err)
	}
}

func getPostHandler(res http.ResponseWriter, req *http.Request) {
	id, err := getIDFromQuery(req)

	if err != nil {
		log.Println(err)
		errorHandler(res, http.StatusNotFound)
		return
	}

	tpl := tmpl.GetTemplateByName("post")
	if err := tpl.Execute(res, data.GetPostByID(id)); err != nil {
		log.Println(err)
	}
}

func editPostHandler(res http.ResponseWriter, req *http.Request) {
	id, err := getIDFromQuery(req)

	if err != nil {
		log.Println(err)
		errorHandler(res, http.StatusNotFound)
		return
	}

	tpl := tmpl.GetTemplateByName("edit")
	if err := tpl.Execute(res, data.GetPostByID(id)); err != nil {
		log.Println(err)
	}
}

func savePostHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		log.Println(err)
		errorHandler(res, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(req.PostFormValue("id"))

	if err != nil {
		log.Println(err)
		errorHandler(res, http.StatusInternalServerError)
		return
	}

	title := req.PostFormValue("title")
	content := req.PostFormValue("content")

	if data.IsPostExists(id) {
		data.UpdatePost(id, title, content)
	} else {
		data.AddNewPost(title, content)
	}

	http.Redirect(res, req, "/", 301)
}

func newPostHandler(res http.ResponseWriter, _ *http.Request) {
	tpl := tmpl.GetTemplateByName("edit")
	if err := tpl.Execute(res, models.NewPost(-1, "", "")); err != nil {
		log.Println(err)
	}
}

func errorHandler(res http.ResponseWriter, status int) {
	res.WriteHeader(status)
	var err error
	switch status {
	case http.StatusBadRequest:
		_, err = fmt.Fprint(res, "bad request")
	case http.StatusNotFound:
		_, err = fmt.Fprint(res, "404 not found")
	case http.StatusInternalServerError:
		_, err = fmt.Fprint(res, "custom 500")
	}
	if err != nil {
		log.Println(err)
	}
}
