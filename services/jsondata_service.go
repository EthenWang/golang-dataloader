package services

import (
	"dataloader/models"
	"dataloader/utils"
	dlerror "dataloader/utils/error"
	"encoding/json"
	"path"
	"reflect"
	"strings"
)

type DataLoaderData interface {
	Prepare()
	Query(id string) interface{}
	// Update(data DataLoaderDataItem) error
	// Delete(id string) error
}

type JsonDataService interface {
	Load(name string) error
	Query(id string) interface{}
}

type jsonDataService struct {
	appConfig *models.AppConfig
	path      string
	typ       reflect.Type
	data      DataLoaderData
}

var cache map[string]map[string]DataLoaderData

func NewJsonDataService(appConfig *models.AppConfig, path string, typ reflect.Type) JsonDataService {
	if cache == nil {
		cache = make(map[string]map[string]DataLoaderData)
	}
	if _, exist := cache[path]; !exist {
		cache[path] = make(map[string]DataLoaderData)
	}

	return &jsonDataService{
		appConfig, path, typ, nil,
	}
}

func (s *jsonDataService) Load(name string) error {
	if _, exist := cache[s.path][name]; exist {
		return nil
	}

	jsonData, err := utils.ReadJson(s.getPath(name))
	if err != nil {
		return err
	}

	pdata := reflect.New(s.typ)
	data := pdata.Elem()
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return err
	}

	dlData, ok := pdata.Interface().(DataLoaderData)
	if !ok {
		return dlerror.NewDataLoaderError(dlerror.InvalidDataType)
	}

	dlData.Prepare()
	s.data = dlData
	cache[s.path][name] = dlData

	return nil
}

func (s *jsonDataService) Query(id string) interface{} {
	if s.data == nil {
		return dlerror.NewDataLoaderError(dlerror.DataIsNil)
	}
	return s.data.Query(id)
}

func (s *jsonDataService) getPath(lang string) string {
	return path.Join(s.appConfig.DataPath, s.path, strings.Join([]string{lang, ".json"}, ""))
}
