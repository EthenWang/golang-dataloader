package models

type MessageModel struct {
	DsMessages MessageDataWraper `json:"ds-messages"`
	_dict      map[string]*MessageItem
}

type MessageDataWraper struct {
	Messages []MessageItem `json:"tt-messages"`
}

type MessageItem struct {
	SystemId      string `json:"system-id"`
	Lang          string `json:"obj-language"`
	Region        string `json:"obj-region"`
	Id            string `json:"message-number"`
	Text          string `json:"message-description"`
	Type          string `json:"message-type"`
	LayerCode     string `json:"layer-code"`
	UserOrGroupId string `json:"user-or-group-id"`
	LayerType     string `json:"layer-type"`
	// UpdateBy           string  `json:"update-by"`
	// UpdateTime         string  `json:"update-time"`
	// UpdateDate         string  `json:"update-date"`
	// CreatedBy          string  `json:"created-by"`
	// CreatedDate        string  `json:"created-date"`
	// ReservedStandard   string  `json:"reserved-standard"`
	// ReservedCustom     string  `json:"reserved-custom"`
	// ChangeReferenceNum float32 `json:"change-reference-num"`
	// ReservedFree       string  `json:"reserved-free"`
	// CreatedTime        string  `json:"created-time"`
	// ReservedSupport    string  `json:"reserved-support"`
}

func (s *MessageModel) New() DataLoaderData {
	return &MessageModel{}
}

func (s *MessageModel) Prepare() {
	if mes := s.DsMessages.Messages; mes != nil {
		s._dict = make(map[string]*MessageItem)
		for i := 0; i < len(mes); i++ {
			s._dict[mes[i].Id] = &mes[i]
		}
	}
}

func (s *MessageModel) All() interface{} {
	return s.DsMessages.Messages
}

func (s *MessageModel) Query(id string) interface{} {
	if tran, exist := s._dict[id]; exist {
		return *tran
	}
	return nil
}
