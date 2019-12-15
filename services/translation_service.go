package services

import (
	"dataloader/models"
	"dataloader/utils"
	dlerror "dataloader/utils/error"
	"encoding/json"
	"path"
	"strings"
)

type translationCollection map[string]*models.TranslationModel

var (
	basePath         string
	translationArray = make(map[models.Language]*[]models.TranslationModel)
	translationMap   = make(map[models.Language]translationCollection)
)

type TranslationService interface {
	GetAll(lang models.Language) ([]models.TranslationModel, error)
	Get(id string, lang models.Language) (models.TranslationModel, error)
	GetList(ids []string, lang models.Language) ([]models.TranslationModel, error)
	Update(tran models.TranslationModel, lang models.Language) error
	UpdateList(tranList []models.TranslationModel, lang models.Language) error
	Delete(id string, lang models.Language) error
	DeleteList(ids []string, lang models.Language) error
	GetAvailableLanguages() ([]models.Language, error)
}

func NewTranslationService(config *models.AppConfig) TranslationService {
	basePath = path.Join(config.DataPath, "sd-translation")
	return &translationService{}
}

type translationService struct{}

func (*translationService) GetAll(lang models.Language) ([]models.TranslationModel, error) {
	err := loadData(lang)
	if err == nil {
		return *translationArray[lang], nil
	}
	return nil, err
}

func (*translationService) Get(id string, lang models.Language) (models.TranslationModel, error) {
	err := loadData(lang)
	if err == nil {
		if tran, ok := translationMap[lang][id]; ok {
			return *tran, nil
		}
	}
	return models.TranslationModel{}, err
}

func (s *translationService) GetList(ids []string, lang models.Language) ([]models.TranslationModel, error) {
	res := make([]models.TranslationModel, 0)
	for _, id := range ids {
		tran, err := s.Get(id, lang)
		if err != nil {
			return nil, err
		}
		res = append(res, tran)
	}
	return res, nil
}

func updateTranslation(tran models.TranslationModel, lang models.Language) error {
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

func (s *translationService) Update(tran models.TranslationModel, lang models.Language) error {
	err := updateTranslation(tran, lang)
	if err != nil {
		return err
	}
	utils.WriteJson(translationArray[lang], getPath(lang))
	return nil
}

func (s *translationService) UpdateList(list []models.TranslationModel, lang models.Language) error {
	if list != nil {
		for _, tran := range list {
			updateTranslation(tran, lang)
		}
		utils.WriteJson(translationArray[lang], getPath(lang))
	}
	return nil
}

func deleteTranslation(id string, lang models.Language) error {
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

func (*translationService) Delete(id string, lang models.Language) error {
	err := deleteTranslation(id, lang)
	if err != nil {
		return err
	}
	utils.WriteJson(translationArray[lang], getPath(lang))
	return nil
}

func (s *translationService) DeleteList(ids []string, lang models.Language) error {
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

func (*translationService) GetAvailableLanguages() ([]models.Language, error) {
	list, err := utils.GetFileList(basePath)
	if err != nil {
		fileList := make([]models.Language, 0)
		for _, l := range list {
			fileList = append(fileList, models.Language(l))
		}
		return fileList, nil
	}
	return nil, err
}

func validate(tran models.TranslationModel, lang models.Language) error {
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

func getPath(lang models.Language) string {
	return path.Join(basePath, strings.Join([]string{string(lang), ".json"}, ""))
}

func loadData(lang models.Language) error {
	if translationMap[lang] == nil {
		data, err := utils.ReadJson(getPath(lang))
		if err != nil {
			return err
		}

		var arr []models.TranslationModel
		err = json.Unmarshal(*data, &arr)
		if err != nil {
			return err
		}

		translationArray[lang] = &arr
		m := make(translationCollection)
		for _, t := range arr {
			m[t.Id] = &t
		}
		translationMap[lang] = m
	}
	return nil
}
