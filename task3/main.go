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
		errorHandler(res, http.StatusNotFound)
		return
	}

	tpl := tpls.GetTemplateByName("index")
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

	tpl := tpls.GetTemplateByName("post")
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

	tpl := tpls.GetTemplateByName("edit")
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

func newPostHandler(res http.ResponseWriter, req *http.Request) {
	tpl := tpls.GetTemplateByName("edit")
	if err := tpl.Execute(res, models.NewPost(-1, "", "")); err != nil {
		log.Println(err)
	}
}

func errorHandler(res http.ResponseWriter, status int) {
	res.WriteHeader(status)
	switch status {
	case http.StatusBadRequest:
		if _, err := fmt.Fprint(res, "bad request"); err != nil {
			log.Println(err)
		}
	case http.StatusNotFound:
		if _, err := fmt.Fprint(res, "404 not found"); err != nil {
			log.Println(err)
		}
	case http.StatusInternalServerError:
		if _, err := fmt.Fprint(res, "custom 500"); err != nil {
			log.Println(err)
		}
	}
}
