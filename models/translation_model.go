package models

import "fmt"

type TranslationModel struct {
	DsTranslation TranslationDataWraper `json:"ds-translation"`
}

type TranslationDataWraper struct {
	Translations []TranslationItem `json:"tt-translation"`
}

type TranslationItem struct {
	Id       string `json:"sd-code"`
	Text     string `json:"sd-text"`
	Lang     string `json:"sd-language"`
	SystemId string `json:"system-id"`
	Region   string `json:"sd-region"`
	// "system-id": "Apprise",
	// "sd-language": "DEFAULT",
	// "sd-region": "",
	// "sd-text": "Additional Charges",
	// "update-by": "system",
	// "update-time": "15:05:11",
	// "update-date": "2005-11-18",
	// "created-by": "system",
	// "created-date": "2005-11-18",
	// "reserved-standard": "",
	// "reserved-custom": "",
	// "sd-code": "00000001",
	// "change-reference-num": 10151259.0,
	// "layer-code": "Default",
	// "user-or-group-id": "",
	// "layer-type": "%DEFAULT%",
	// "reserved-free": "",
	// "created-time": "",
	// "reserved-support": "",
	// "sd-abbreviation": ""
}

func (s *TranslationModel) Id() string {
	return "default"
}

func (s *TranslationModel) Query(id string) *TranslationItem {
	fmt.Println("TranslationModel Query: %s", id)
	return nil
}
