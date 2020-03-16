package controllers

import (
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewPost struct {
	beego.Controller
	Db *mongo.Client
}

func (c *NewPost) Get() {
	c.TplName = "new.tpl"
}
