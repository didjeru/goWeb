package controllers

import (
	"../models"
	"database/sql"
	"github.com/astaxie/beego"
	"strconv"
)

type EditPost struct {
	beego.Controller
	Db *sql.DB
}

func (c *EditPost) Post() {

	idStr := c.Ctx.Request.FormValue("id")
	id, err := strconv.Atoi(idStr)

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

	if err := editPost(c.Db, &post, idStr); err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Redirect("/", 301)
}
