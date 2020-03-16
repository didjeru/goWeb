package controllers

import (
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
)

type Prepare struct {
	beego.Controller
	Db *mongo.Client
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
