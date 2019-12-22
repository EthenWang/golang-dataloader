package controllers

import (
	"dataloader/models"
	"dataloader/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type MessageController struct {
	Ctx iris.Context
	//TranslationService services.TranslationService
	DataService services.JsonDataService
}

func (c *MessageController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/{lang}", "Get")
	b.Handle("GET", "/{lang}/{id}", "GetById")
}

// Get - return all translations
func (c *MessageController) Get(lang string) []models.MessageItem {
	c.DataService.Load(lang)
	data, ok := c.DataService.All().([]models.MessageItem)
	if !ok {
		return nil
	}
	return data
}

// GetById - return a specific translation. param "id" is translation code
func (c *MessageController) GetById(lang string, id string) models.MessageItem {
	c.DataService.Load(lang)
	data, ok := c.DataService.Query(id).(models.MessageItem)
	if !ok {
		return models.MessageItem{}
	}
	return data
}
