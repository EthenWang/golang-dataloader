package services

import (
	"dataloader/models"
	"dataloader/utils"
	"encoding/json"
	"path"
	"strings"
)

type JsonDataService interface {
	Load(name string) (models.DataLoaderData, error)
}

type jsonDataService struct {
	appConfig *models.AppConfig
	path      string
	creator   models.ModelCreator
	data      models.DataLoaderData
}

var cache map[string]map[string]models.DataLoaderData

func NewJsonDataService(appConfig *models.AppConfig, path string, creator models.ModelCreator) JsonDataService {
	if cache == nil {
		cache = make(map[string]map[string]models.DataLoaderData)
	}
	if _, exist := cache[path]; !exist {
		cache[path] = make(map[string]models.DataLoaderData)
	}

	return &jsonDataService{
		appConfig, path, creator, creator.New(),
	}
}

func (s *jsonDataService) Load(name string) (models.DataLoaderData, error) {
	if data, exist := cache[s.path][name]; exist {
		return data, nil
	}

	jsonData, err := utils.ReadJson(s.getPath(name))
	if err != nil {
		return nil, err
	}

	// research how to use reflect to create model dynamically
	// pdata := reflect.New(s.typ)
	// data := pdata.Elem().Interface()

	// fmt.Println("data type", reflect.TypeOf(data))
	// err = json.Unmarshal(jsonData, &data)
	// fmt.Println("data type", reflect.TypeOf(data))
	// fmt.Println(data)
	// if err != nil {
	// 	return err
	// }
	// dlData, ok := pdata.Interface().(DataLoaderData)
	// if !ok {
	// 	return dlerror.NewDataLoaderError(dlerror.InvalidDataType)
	// }

	err = json.Unmarshal(jsonData, s.data)

	if err != nil {
		return nil, err
	}

	if qdata, ok := s.data.(models.Query); ok {
		qdata.Prepare()
	}

	if s.appConfig.UseCache {
		cache[s.path][name] = s.data
	}

	return s.data, nil
}

func (s *jsonDataService) getPath(name string) string {
	return path.Join(s.appConfig.DataPath, s.path, strings.Join([]string{name, ".json"}, ""))
}
