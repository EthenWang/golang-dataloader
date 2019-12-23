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
	data, err := c.DataService.Load(lang)
	if err != nil {
		return nil
	}
	qdata, ok := data.(models.Query)
	if !ok {
		return nil
	}

	return qdata.All().([]models.TranslationItem)
}

// GetById - return a specific translation. param "id" is translation code
func (c *TranslationController) GetById(lang string, id string) models.TranslationItem {
	data, err := c.DataService.Load(lang)
	if err != nil {
		return models.TranslationItem{}
	}
	qdata, ok := data.(models.Query)
	if !ok {
		return models.TranslationItem{}
	}
	return qdata.Query(id).(models.TranslationItem)
}
