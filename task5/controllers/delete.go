package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
)

type DeletePost struct {
	beego.Controller
	Db *sql.DB
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
