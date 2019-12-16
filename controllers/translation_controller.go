package controllers

import (
	"dataloader/models"
	"dataloader/services"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type TranslationController struct {
	Ctx                iris.Context
	TranslationService services.TranslationService
}

func (c *TranslationController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/{lang}", "Get")
	b.Handle("GET", "/{lang}/{id}", "GetById")
}

// Get - return all translations
func (c *TranslationController) Get(lang string) []models.TranslationItem {
	data, _ := c.TranslationService.GetAll(lang)
	return data
}

// GetById - return id
func (c *TranslationController) GetById(lang string, id string) models.TranslationItem {
	data, _ := c.TranslationService.Get(id, lang)
	return data
}
