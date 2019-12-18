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

type translationCollection map[string]*models.TranslationItem

type TranslationService interface {
	GetAll(lang string) ([]models.TranslationItem, error)
	Get(id string, lang string) (models.TranslationItem, error)
	GetList(ids []string, lang string) ([]models.TranslationItem, error)
	Update(tran models.TranslationItem, lang string) error
	UpdateList(tranList []models.TranslationItem, lang string) error
	Delete(id string, lang string) error
	DeleteList(ids []string, lang string) error
	GetAvailableLanguages() ([]string, error)
}

type translationService struct {
	appConfig        *models.AppConfig
	translationMap   map[string]translationCollection
	translationArray map[string]*[]models.TranslationItem
}

func NewTranslationService(config *models.AppConfig) TranslationService {
	return &translationService{
		appConfig:        config,
		translationMap:   make(map[string]translationCollection),
		translationArray: make(map[string]*[]models.TranslationItem),
	}
}

func (s *translationService) GetAll(lang string) ([]models.TranslationItem, error) {
	err := s.loadData(lang)
	if err == nil {
		return *s.translationArray[lang], nil
	}
	return nil, err
}

func (s *translationService) Get(id string, lang string) (models.TranslationItem, error) {
	ds := NewJsonDataService(s.appConfig, reflect.TypeOf(models.TranslationModel{}))
	ds.Query("")
	err := s.loadData(lang)
	if err == nil {
		if tran, ok := s.translationMap[lang][id]; ok {
			return *tran, nil
		}
	}
	return models.TranslationItem{}, err
}

func (s *translationService) GetList(ids []string, lang string) ([]models.TranslationItem, error) {
	res := make([]models.TranslationItem, 0)
	for _, id := range ids {
		tran, err := s.Get(id, lang)
		if err != nil {
			return nil, err
		}
		res = append(res, tran)
	}
	return res, nil
}

func (s *translationService) updateTranslation(tran models.TranslationItem, lang string) error {
	err := s.loadData(lang)
	if err != nil {
		return err
	}

	if trans, ok := s.translationMap[lang]; ok {
		err = s.validate(tran, lang)
		if err != nil {
			return err
		}
		trans[tran.Id] = &tran
		return nil
	}
	return dlerror.NewDataLoaderError(dlerror.LanguageNotExist)
}

func (s *translationService) Update(tran models.TranslationItem, lang string) error {
	err := s.updateTranslation(tran, lang)
	if err != nil {
		return err
	}
	s.saveData(lang)
	return nil
}

func (s *translationService) UpdateList(list []models.TranslationItem, lang string) error {
	if list != nil {
		for _, tran := range list {
			s.updateTranslation(tran, lang)
		}
		s.saveData(lang)
	}
	return nil
}

func (s *translationService) deleteTranslation(id string, lang string) error {
	err := s.loadData(lang)
	if err != nil {
		return err
	}

	if trans, ok := s.translationMap[lang]; ok {
		if _, ok = trans[id]; ok {
			delete(trans, id)
			return nil
		}
	}
	return dlerror.NewDataLoaderError(dlerror.LanguageNotExist)
}

func (s *translationService) Delete(id string, lang string) error {
	err := s.deleteTranslation(id, lang)
	if err != nil {
		return err
	}
	s.saveData(lang)
	return nil
}

func (s *translationService) DeleteList(ids []string, lang string) error {
	if ids != nil {
		for _, id := range ids {
			err := s.deleteTranslation(id, lang)
			if err != nil {
				return err
			}
		}
		s.saveData(lang)
	}
	return nil
}

func (s *translationService) GetAvailableLanguages() ([]string, error) {
	list, err := utils.GetFileList(s.appConfig.DataPath)
	if err != nil {
		fileList := make([]string, 0)
		for _, l := range list {
			fileList = append(fileList, string(l))
		}
		return fileList, nil
	}
	return nil, err
}

func (s *translationService) validate(tran models.TranslationItem, lang string) error {
	if trans, ok := s.translationArray[lang]; ok {
		for _, t := range *trans {
			if t.Text == tran.Text {
				return dlerror.NewDataLoaderError(dlerror.TranslationExist)
			}
		}
		return nil
	}
	return dlerror.NewDataLoaderError(dlerror.LanguageNotExist)
}

func (s *translationService) getPath(lang string) string {
	return path.Join(s.appConfig.DataPath, strings.Join([]string{lang, ".json"}, ""))
}

func (s *translationService) loadData(lang string) error {
	if s.translationMap[lang] == nil {
		jsonData, err := utils.ReadJson(s.getPath(lang))
		if err != nil {
			return err
		}

		var transData models.TranslationModel
		err = json.Unmarshal(*jsonData, &transData)
		if err != nil {
			return err
		}

		arr := transData.DsTranslation.Translations
		s.translationArray[lang] = &arr
		m := make(translationCollection)
		for i := 0; i < len(arr); i++ {
			m[arr[i].Id] = &arr[i]
		}
		s.translationMap[lang] = m
	}
	return nil
}

func (s *translationService) saveData(lang string) error {
	data := models.TranslationModel{
		DsTranslation: models.TranslationDataWraper{
			Translations: *s.translationArray[lang],
		},
	}
	return utils.WriteJson(data, s.getPath(lang))
}
