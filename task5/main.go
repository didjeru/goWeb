package main

import (
	_ "./routers"
	"github.com/astaxie/beego"
	"os"
)

func main() {
	beego.Run("localhost", os.Getenv("httpport"))
}
