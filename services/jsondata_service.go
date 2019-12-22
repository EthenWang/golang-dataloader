package services

import (
	"dataloader/models"
	"dataloader/utils"
	dlerror "dataloader/utils/error"
	"encoding/json"
	"path"
	"strings"
)

type JsonDataService interface {
	Load(name string) error
	All() interface{}
	Query(id string) interface{}
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

func (s *jsonDataService) Load(name string) error {
	if _, exist := cache[s.path][name]; exist {
		return nil
	}

	jsonData, err := utils.ReadJson(s.getPath(name))
	if err != nil {
		return err
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
		return err
	}

	s.data.Prepare()
	if s.appConfig.UseCache {
		cache[s.path][name] = s.data
	}

	return nil
}

func (s *jsonDataService) All() interface{} {
	if s.data == nil {
		return dlerror.NewDataLoaderError(dlerror.DataIsNil)
	}
	return s.data.All()
}

func (s *jsonDataService) Query(id string) interface{} {
	if s.data == nil {
		return dlerror.NewDataLoaderError(dlerror.DataIsNil)
	}
	return s.data.Query(id)
}

func (s *jsonDataService) getPath(name string) string {
	return path.Join(s.appConfig.DataPath, s.path, strings.Join([]string{name, ".json"}, ""))
}
