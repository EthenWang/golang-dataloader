package models

type TranslationModel struct {
	DsTranslation struct {
		Translations []TranslationItem `json:"tt-translation"`
	} `json:"ds-translation"`
	_dict map[string]*TranslationItem
}

type TranslationItem struct {
	SystemId       string `json:"system-id"`
	Lang           string `json:"sd-language"`
	Region         string `json:"sd-region"`
	Id             string `json:"sd-code"`
	Text           string `json:"sd-text"`
	LayerCode      string `json:"layer-code"`
	UserOrGroupId  string `json:"user-or-group-id"`
	LayerType      string `json:"layer-type"`
	SdAbbreviation string `json:"sd-abbreviation"`
}

type translationJsonItem struct {
	*TranslationItem
	UpdateBy           string  `json:"update-by"`
	UpdateTime         string  `json:"update-time"`
	UpdateDate         string  `json:"update-date"`
	CreatedBy          string  `json:"created-by"`
	CreatedDate        string  `json:"created-date"`
	ReservedStandard   string  `json:"reserved-standard"`
	ReservedCustom     string  `json:"reserved-custom"`
	ChangeReferenceNum float32 `json:"change-reference-num"`
	ReservedFree       string  `json:"reserved-free"`
	CreatedTime        string  `json:"created-time"`
	ReservedSupport    string  `json:"reserved-support"`
}

func (s *TranslationModel) New() DataLoaderData {
	return &TranslationModel{}
}

func (s *TranslationModel) Prepare() {
	if trans := s.DsTranslation.Translations; trans != nil {
		s._dict = make(map[string]*TranslationItem)
		for i := 0; i < len(trans); i++ {
			s._dict[trans[i].Id] = &trans[i]
		}
	}
}

func (s *TranslationModel) All() interface{} {
	return s.DsTranslation.Translations
}

func (s *TranslationModel) Query(id string) interface{} {
	if tran, exist := s._dict[id]; exist {
		return *tran
	}
	return nil
}
