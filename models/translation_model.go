package models

type Language string

//type TranslationModel struct {
//	Translations []TranslationItem `json:"tt-translation"`
//}

type TranslationModel struct {
	Id   string `json:"sd-code"`
	Text string `json:"sd-text"`
	Lang Language
}
