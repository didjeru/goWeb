package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
)

type NewPost struct {
	beego.Controller
	Db *sql.DB
}

func (c *NewPost) Get() {
	c.TplName = "new.tpl"
}
