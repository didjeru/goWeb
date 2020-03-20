package controllers

import (
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
)

type MainController struct {
	beego.Controller
	Db *mongo.Client
}

func (c *MainController) Get() {
	posts, err := getPosts(c.Db)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(500)
		_, _ = c.Ctx.ResponseWriter.Write([]byte(err.Error()))
		return
	}

	c.Data["Posts"] = posts
	c.TplName = "index.tpl"
}
