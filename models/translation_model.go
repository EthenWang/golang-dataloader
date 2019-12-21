package models

type TranslationModel struct {
	DsTranslation TranslationDataWraper `json:"ds-translation"`
	_dict         map[string]*TranslationItem
}

type TranslationDataWraper struct {
	Translations []TranslationItem `json:"tt-translation"`
}

type TranslationItem struct {
	SystemId string `json:"system-id"`
	Lang     string `json:"sd-language"`
	Region   string `json:"sd-region"`
	Id       string `json:"sd-code"`
	Text     string `json:"sd-text"`
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

func (s *TranslationModel) Prepare() {
	if trans := s.DsTranslation.Translations; trans != nil {
		s._dict = make(map[string]*TranslationItem)
		for i := 0; i < len(trans); i++ {
			s._dict[trans[i].Id] = &trans[i]
		}
	}
}

func (s *TranslationModel) Query(id string) interface{} {
	if tran, exist := s._dict[id]; exist {
		return tran
	}
	return nil
}
