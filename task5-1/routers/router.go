package routers

import (
	"../controllers"
	"../models"
	"database/sql"
	"github.com/astaxie/beego"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DSN = "root:qwerty@tcp(localhost:3306)/blog?charset=utf8"
)

type DB struct {
	sql   *sql.DB
	posts map[int]models.Post
}

func init() {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	beego.Router("/", &controllers.MainController{
		Db: db,
	})

	beego.Router("/post/:id", &controllers.SinglePost{
		Db: db,
	})

	beego.Router("/prepare/:id", &controllers.Prepare{
		Db: db,
	})
}
