package services

import (
	"dataloader/models"
	"dataloader/utils"
	dlerror "dataloader/utils/error"
	"encoding/json"
	"path"
	"strings"
)

type translationCollection map[string]*models.TranslationItem

var (
	basePath         string
	translationArray = make(map[string]*[]models.TranslationItem)
	translationMap   = make(map[string]translationCollection)
)

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

func NewTranslationService(config *models.AppConfig) TranslationService {
	basePath = path.Join(config.DataPath, "sd-translation")
	return &translationService{}
}

type translationService struct{}

func (*translationService) GetAll(lang string) ([]models.TranslationItem, error) {
	err := loadData(lang)
	if err == nil {
		return *translationArray[lang], nil
	}
	return nil, err
}

func (*translationService) Get(id string, lang string) (models.TranslationItem, error) {
	err := loadData(lang)
	if err == nil {
		if tran, ok := translationMap[lang][id]; ok {
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

func updateTranslation(tran models.TranslationItem, lang string) error {
	err := loadData(lang)
	if err != nil {
		return err
	}

	if trans, ok := translationMap[lang]; ok {
		err = validate(tran, lang)
		if err != nil {
			return err
		}
		trans[tran.Id] = &tran
		return nil
	}
	return dlerror.NewDataLoaderError(dlerror.LanguageNotExist)
}

func (s *translationService) Update(tran models.TranslationItem, lang string) error {
	err := updateTranslation(tran, lang)
	if err != nil {
		return err
	}
	utils.WriteJson(translationArray[lang], getPath(lang))
	return nil
}

func (s *translationService) UpdateList(list []models.TranslationItem, lang string) error {
	if list != nil {
		for _, tran := range list {
			updateTranslation(tran, lang)
		}
		utils.WriteJson(translationArray[lang], getPath(lang))
	}
	return nil
}

func deleteTranslation(id string, lang string) error {
	err := loadData(lang)
	if err != nil {
		return err
	}

	if trans, ok := translationMap[lang]; ok {
		if _, ok = trans[id]; ok {
			delete(trans, id)
			return nil
		}
	}
	return dlerror.NewDataLoaderError(dlerror.LanguageNotExist)
}

func (*translationService) Delete(id string, lang string) error {
	err := deleteTranslation(id, lang)
	if err != nil {
		return err
	}
	utils.WriteJson(translationArray[lang], getPath(lang))
	return nil
}

func (s *translationService) DeleteList(ids []string, lang string) error {
	if ids != nil {
		for _, id := range ids {
			err := deleteTranslation(id, lang)
			if err != nil {
				return err
			}
		}
		utils.WriteJson(translationArray[lang], getPath(lang))
	}
	return nil
}

func (*translationService) GetAvailableLanguages() ([]string, error) {
	list, err := utils.GetFileList(basePath)
	if err != nil {
		fileList := make([]string, 0)
		for _, l := range list {
			fileList = append(fileList, string(l))
		}
		return fileList, nil
	}
	return nil, err
}

func validate(tran models.TranslationItem, lang string) error {
	if trans, ok := translationArray[lang]; ok {
		for _, t := range *trans {
			if t.Text == tran.Text {
				return dlerror.NewDataLoaderError(dlerror.TranslationExist)
			}
		}
		return nil
	}
	return dlerror.NewDataLoaderError(dlerror.LanguageNotExist)
}

func getPath(lang string) string {
	return path.Join(basePath, strings.Join([]string{string(lang), ".json"}, ""))
}

func loadData(lang string) error {
	if translationMap[lang] == nil {
		jsonData, err := utils.ReadJson(getPath(lang))
		if err != nil {
			return err
		}

		var transData models.TranslationModel
		err = json.Unmarshal(*jsonData, &transData)
		if err != nil {
			return err
		}

		translationArray[lang] = &transData.DsTranslation.Translations
		m := make(translationCollection)
		for _, t := range *translationArray[lang] {
			m[t.Id] = &t
		}
		translationMap[lang] = m
	}
	return nil
}
