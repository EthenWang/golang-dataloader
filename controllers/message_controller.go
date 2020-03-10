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

// Get - return all messages
func (c *MessageController) Get(lang string) []models.MessageItem {
	data, err := c.DataService.Load(lang)
	if err != nil {
		return nil
	}
	qdata, ok := data.(models.Query)
	if !ok {
		return nil
	}

	return qdata.All().([]models.MessageItem)
}

// GetById - return a specific message. param "id" is message code
func (c *MessageController) GetById(lang string, id string) models.MessageItem {
	data, err := c.DataService.Load(lang)
	if err != nil {
		return models.MessageItem{}
	}
	qdata, ok := data.(models.Query)
	if !ok {
		return models.MessageItem{}
	}
	return qdata.Query(id).(models.MessageItem)
}
