package controllers

import (
	"../models"
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
)

type SinglePost struct {
	beego.Controller
	Db *mongo.Client
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

	post := models.Post{
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
