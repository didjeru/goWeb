package routers

import (
	"../controllers"
	"context"
	"github.com/astaxie/beego"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client

func init() {
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	client = db
	err = client.Connect(context.Background())
	if err != nil {
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

	beego.Router("/edit/:id", &controllers.EditPost{
		Db: db,
	})

	beego.Router("/new", &controllers.NewPost{
		Db: db,
	})

	beego.Router("/post", &controllers.SinglePost{
		Db: db,
	})

	beego.Router("/delete/:id", &controllers.DeletePost{
		Db: db,
	})

}
