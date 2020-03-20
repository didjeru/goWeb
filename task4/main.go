package main

import (
	"./db"
	"./models"
	"./templates"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var data = db.Init()
var tpls = templates.Init()

func main() {
	http.HandleFunc("/new", newPostHandler)
	http.HandleFunc("/edit", editPostHandler)
	http.HandleFunc("/post", getPostHandler)
	http.HandleFunc("/save", savePostHandler)
	http.HandleFunc("/", indexHandler)

	port := "8080"
	log.Println("Server started at port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getIDFromQuery(req *http.Request) (int, error) {
	keys, ok := req.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		return -1, errors.New("Url Param 'id' is missing")
	}

	return strconv.Atoi(keys[0])
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		errorHandler(res, req, http.StatusNotFound)
		return
	}

	tpl := tpls.GetTemplateByName("index")
	tpl.Execute(res, data.GetAllPosts())
}

func getPostHandler(res http.ResponseWriter, req *http.Request) {
	id, err := getIDFromQuery(req)

	if err != nil {
		log.Println(err)
		errorHandler(res, req, http.StatusNotFound)
		return
	}

	tpl := tpls.GetTemplateByName("post")
	tpl.Execute(res, data.GetPostByID(id))
}

func editPostHandler(res http.ResponseWriter, req *http.Request) {
	id, err := getIDFromQuery(req)

	if err != nil {
		log.Println(err)
		errorHandler(res, req, http.StatusNotFound)
		return
	}

	tpl := tpls.GetTemplateByName("edit")
	tpl.Execute(res, data.GetPostByID(id))
}

func savePostHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		log.Println(err)
		errorHandler(res, req, http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(req.PostFormValue("id"))

	if err != nil {
		log.Println(err)
		errorHandler(res, req, http.StatusInternalServerError)
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

func newPostHandler(res http.ResponseWriter, req *http.Request) {
	tpl := tpls.GetTemplateByName("edit")
	tpl.Execute(res, models.NewPost(-1, "", ""))
}

func errorHandler(res http.ResponseWriter, req *http.Request, status int) {
	res.WriteHeader(status)

	if status == http.StatusBadRequest {
		fmt.Fprint(res, "bad request")
	}

	if status == http.StatusNotFound {
		fmt.Fprint(res, "custom 404")
	}

	if status == http.StatusInternalServerError {
		fmt.Fprint(res, "custom 500")
	}
}
