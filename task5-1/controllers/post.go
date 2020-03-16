package controllers

import (
	"../models"
	"database/sql"
	"encoding/json"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"strconv"
)

type SinglePost struct {
	beego.Controller
	Db *sql.DB
}

func (c *SinglePost) Get() {
	post, err := getPost(c.Db, c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Data["Post"] = post
	c.TplName = "single.tpl"
}

func (c *SinglePost) Post() {
	id, err := strconv.Atoi(c.Ctx.Request.FormValue("id"))
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	post := models.Post{
		ID:      id,
		Title:   c.Ctx.Request.FormValue("title"),
		Content: c.Ctx.Request.FormValue("content"),
	}

	if err := addPost(c.Db, post); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Redirect("/", 301)
}

func (c *SinglePost) Put() {
	id := c.Ctx.Input.Param(":id")
	if len(id) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte("empty id"))
		return
	}

	post := new(models.Post)
	if err := readAndUnmarshall(post, c.Ctx.Request.Body); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	if err := editPost(c.Db, post, id); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Redirect("/", 301)
}

func readAndUnmarshall(resp interface{}, body io.ReadCloser) error {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, resp); err != nil {
		return err
	}

	return nil
}

func (c *SinglePost) Delete() {
	id := c.Ctx.Input.Param(":id")
	if len(id) == 0 {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte("empty id"))
		return
	}

	if err := deletePost(c.Db, id); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Redirect("/", 301)
}
