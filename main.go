package main

import (
	"dataloader/controllers"
	"dataloader/models"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var appconfig models.AppConfig

func main() {
	app := iris.New()
	//app.Logger().SetLevel("debug")
	mvc.Configure(app.Party("/translation/{id}"), translation)
	mvc.Configure(app.Party("/messages/{id}"), messages)
	mvc.Configure(app.Party("/screen/{id}"), screen)

	fmt.Println("app started")
	app.Run(
		iris.Addr("localhost:7500"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

func translation(app *mvc.Application) {
	app.Handle(new(controllers.TranslationController))
}

func messages(app *mvc.Application) {

}

func screen(app *mvc.Application) {

}
