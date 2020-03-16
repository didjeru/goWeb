package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
)

type Prepare struct {
	beego.Controller
	Db *sql.DB
}

func (c *Prepare) Get() {
	post, err := getPost(c.Db, c.Ctx.Input.Param(":id"))
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Data["Post"] = post
	c.TplName = "edit.tpl"
}
