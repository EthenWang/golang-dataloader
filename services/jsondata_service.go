package services

import (
	"dataloader/models"
	"dataloader/utils"
	"encoding/json"
	"path"
	"reflect"
	"strings"
)

type JsonDataService interface {
	Query(id string) interface{}
	// Update(data DataLoaderDataItem) error
	// Delete(id string) error
}

type DataLoaderData interface {
	Id() string
	JsonDataService
}

type jsonDataService struct {
	appConfig *models.AppConfig
	data      reflect.Value
}

func NewJsonDataService(appConfig *models.AppConfig, typ reflect.Type) JsonDataService {
	return &jsonDataService{
		appConfig, reflect.New(typ),
	}
}

func (s *jsonDataService) Query(id string) interface{} {
	jsonData, err := utils.ReadJson(s.getPath("sd-translation\\default"))
	if err != nil {
		return nil
	}

	data := s.data.Elem()
	err = json.Unmarshal(*jsonData, &data)
	if err != nil {
		return nil
	}

	if dlData, ok := s.data.Interface().(DataLoaderData); ok {
		return dlData.Query(id)
	}

	return nil
}

func (s *jsonDataService) getPath(lang string) string {
	return path.Join(s.appConfig.DataPath, strings.Join([]string{lang, ".json"}, ""))
}
