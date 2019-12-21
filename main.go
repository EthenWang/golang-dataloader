package main

import (
	"dataloader/controllers"
	"dataloader/models"
	"dataloader/services"
	"dataloader/utils"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

var appconfig models.AppConfig

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	readConfig()

	mvc.Configure(app.Party("/translation"), translation)
	mvc.Configure(app.Party("/messages"), messages)
	mvc.Configure(app.Party("/screen"), screen)

	app.Run(
		iris.Addr("localhost:7500"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

func translation(app *mvc.Application) {
	//app.Register(services.NewTranslationService(&appconfig))
	app.Register(services.NewJsonDataService(
		&appconfig,
		"sd-translation",
		reflect.TypeOf(models.TranslationModel{}),
	))
	app.Handle(new(controllers.TranslationController))
}

func messages(app *mvc.Application) {

}

func screen(app *mvc.Application) {

}

func readConfig() {
	data, err := utils.ReadJson("app.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &appconfig)
	if err != nil {
		fmt.Println(err)
	}
}
