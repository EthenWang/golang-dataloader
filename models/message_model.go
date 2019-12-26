package models

import "dataloader/utils"

type MessageModel struct {
	DsMessages struct {
		Messages []MessageItem `json:"tt-messages"`
	} `json:"ds-messages"`
	_dict map[string]*MessageItem
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
}

type messageJsonModel struct {
	DsMessages messageJsonItemWrapper `json:"ds-messages"`
}

type messageJsonItemWrapper struct {
	Messages []messageJsonItem `json:"tt-messages"`
}

type messageJsonItem struct {
	*MessageItem
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

func (m *MessageModel) New() DataLoaderData {
	return &MessageModel{}
}

func (m *MessageModel) Prepare() {
	if mes := m.DsMessages.Messages; mes != nil {
		m._dict = make(map[string]*MessageItem)
		for i := 0; i < len(mes); i++ {
			m._dict[mes[i].Id] = &mes[i]
		}
	}
}

func (m *MessageModel) All() interface{} {
	return m.DsMessages.Messages
}

func (m *MessageModel) Query(id string) interface{} {
	if tran, exist := m._dict[id]; exist {
		return *tran
	}
	return nil
}

func (m *MessageModel) Save(path string) error {
	var data = make([]messageJsonItem, 0)
	for i := 0; i < len(m.DsMessages.Messages); i++ {
		data = append(data, messageJsonItem{
			MessageItem:        &m.DsMessages.Messages[i],
			UpdateBy:           "",
			UpdateTime:         "",
			UpdateDate:         "",
			CreatedBy:          "",
			CreatedDate:        "",
			ReservedStandard:   "",
			ReservedCustom:     "",
			ChangeReferenceNum: 0,
			ReservedFree:       "",
			CreatedTime:        "",
			ReservedSupport:    "",
		})
	}
	return utils.WriteJson(messageJsonModel{
		DsMessages: messageJsonItemWrapper{
			Messages: data,
		},
	}, path)
}
