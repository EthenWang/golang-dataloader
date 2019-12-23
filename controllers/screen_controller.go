package controllers

import (
	"dataloader/models"
	"dataloader/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type ScreenController struct {
	Ctx iris.Context
	//TranslationService services.TranslationService
	DataService services.JsonDataService
}

func (c *ScreenController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/{id}", "GetById")
	//b.Handle("GET", "/list", "GetScreenList")
}

// GetById - return a specific translation. param "id" is translation code
func (c *ScreenController) GetById(id string) models.ScreenDataModel {
	data, err := c.DataService.Load(id)
	if err != nil {
		return models.ScreenDataModel{}
	}
	return *data.(*models.ScreenDataModel)
}
