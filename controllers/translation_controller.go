package controllers

import (
	"dataloader/models"
	"dataloader/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type TranslationController struct {
	Ctx iris.Context
	//TranslationService services.TranslationService
	DataService services.JsonDataService
}

func (c *TranslationController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/{lang}", "Get")
	b.Handle("GET", "/{lang}/{id}", "GetById")
}

// Get - return all translations
func (c *TranslationController) Get(lang string) []models.TranslationItem {
	//data, _ := c.TranslationService.GetAll(lang)
	//return data
	return nil
}

// GetById - return a specific translation. param "id" is translation code
func (c *TranslationController) GetById(lang string, id string) models.TranslationItem {
	c.DataService.Load(lang)
	data, ok := c.DataService.Query(id).(models.TranslationItem)
	if !ok {
		return models.TranslationItem{}
	}
	return data
}
