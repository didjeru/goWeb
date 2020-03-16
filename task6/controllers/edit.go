package controllers

import (
	"../models"
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
)

type EditPost struct {
	beego.Controller
	Db *mongo.Client
}

func (c *EditPost) Post() {

	id := c.Ctx.Input.Param(":id")

	post := models.Post{
		ID:      id,
		Title:   c.Ctx.Request.FormValue("title"),
		Content: c.Ctx.Request.FormValue("content"),
	}

	if err := editPost(c.Db, &post, id); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Redirect("/", 301)
}
