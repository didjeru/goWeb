package controllers

import (
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeletePost struct {
	beego.Controller
	Db *mongo.Client
}

func (c *DeletePost) Get() {
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
