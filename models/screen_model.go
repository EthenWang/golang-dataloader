package models

type ScreenDataModel struct {
	DsScreen struct {
		ScreenDef      []ScreenModel         `json:"tt-screen"`
		ObjectDef      []ScreenObjModel      `json:"tt-screen-obj"`
		ChildObjectDef []ScreenChildObjModel `json:"tt-screen-child-obj"`
	} `json:"ds-screen"`
}

type ScreenModel struct {
	SystemId      string `json:"system-id"`
	LayerCode     string `json:"layer-code"`
	UserOrGroupId string `json:"user-or-group-id"`
	LayerType     string `json:"layer-type"`
	Name          string `json:"screen-name"`
	Program       string `json:"screen-program"`
	ScreenType    string `json:"screen-type"`
	TemplateCode  string `json:"template-code"`
	SpecialCase   string `json:"special-case"`
	HelpKey       string `json:"help-key"`
	FocusField    string `json:"focus-field"`
	Title         string `json:"screen-title"`
	ToolTip       string `json:"screen-tooltip"`
	VistualDates  bool   `json:"virtual-dates"`
	// "maint-list-type": "None",
	// "maint-drop-name": "",
	// "mid-update": "",
	// "auto-scr-sync": false
	// "maint-list-obj": "",
	// "skip-layer": false,
	// "change-reference-num": 0.0,
	// "reserved-free": "",
	// "created-time": "",
	// "reserved-support": "",
	// "update-by": "ethen.wang",
	// "update-time": "21:14:49",
	// "update-date": "2019-12-20",
	// "created-by": "ScreenConvert",
	// "created-date": "2019-04-12",
	// "reserved-standard": "test",
	// "reserved-custom": "",
}

type ScreenObjModel struct {
	SystemId          string  `json:"system-id"`
	LayerCode         string  `json:"layer-code"`
	UserOrGroupId     string  `json:"user-or-group-id"`
	LayerType         string  `json:"layer-type"`
	Name              string  `json:"screen-obj-name"`
	Type              string  `json:"obj-type"`
	Format            string  `json:"obj-format"`
	FormatType        string  `json:"obj-format-type"`
	ScreenName        string  `json:"screen-name"`
	Sequence          int     `json:"screen-order"`
	LabelCode         string  `json:"obj-label-code"`
	ToolTip           string  `json:"obj-tooltip-code"`
	IsActive          bool    `json:"obj-active"`
	IsRequired        bool    `json:"obj-required"`
	IsVisible         bool    `json:"obj-visible"`
	IsEnabled         bool    `json:"obj-enabled"`
	IsDisplay         bool    `json:"obj-display"`
	IsUnLabelled      bool    `json:"obj-unlabelled"`
	IsNeedValidation  bool    `json:"obj-needs-validation"`
	Parent            string  `json:"obj-parent"`
	Row               float32 `json:"obj-row"`
	RowSpan           int     `json:"row-span"`
	Column            float32 `json:"obj-col"`
	ColumnSpan        int     `json:"col-span"`
	Height            float32 `json:"obj-height"`
	Width             float32 `json:"obj-width"`
	Image             string  `json:"obj-image"`
	SpecialCase       string  `json:"special-case"`
	InitValue         string  `json:"obj-initial-value"`
	EventList         string  `json:"obj-process-response"`
	Align             string  `json:"obj-align"`
	CellAlign         string  `json:"obj-cell-align"`
	GridAttributes    string  `json:"grid-attributes"`
	UserDefinedType   string  `json:"user-def-type"`
	UserDefinedNumber int     `json:"user-def-num"`
	RelatedTabProgram string  `json:"rel-tabpage-pgm"`
	DropName          string  `json:"drop-name"`
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
	// "linked-browse-has-params": false,
	// "browse-filter-num-1": 0,
	// "browse-param-type-1": "",
	// "browse-param-value-1": "",
	// "browse-filter-num-2": 0,
	// "browse-param-type-2": "",
	// "browse-param-value-2": "",
	// "browse-filter-num-3": 0,
	// "browse-param-type-3": "",
	// "browse-param-value-3": "",
	// "obj-colon": false,
	// "button-tooltip-code": "",
	// "hotkey-code": "",
	// "hotkey-hdr-name": "",
	// "link-collection-parent": "",
	// "always-hidden": false,
	// "always-disabled": false,
	// "supports-dot-notation": false,
	// "no-related-table": true,
	// "table-name": "",
	// "field-name": "",
	// "help-override-single-language": "",
	// "field-type": "",
	// "obj-created-by-screen-designer": false
}

type ScreenChildObjModel struct {
	SystemId          string  `json:"system-id"`
	LayerCode         string  `json:"layer-code"`
	UserOrGroupId     string  `json:"user-or-group-id"`
	LayerType         string  `json:"layer-type"`
	Name              string  `json:"obj-name"`
	Type              string  `json:"obj-type"`
	Format            string  `json:"obj-format"`
	FormatType        string  `json:"obj-format-type"`
	ScreenName        string  `json:"screen-name"`
	Sequence          int     `json:"screen-order"`
	LabelCode         string  `json:"obj-label-code"`
	ToolTip           string  `json:"obj-tooltip-code"`
	Parent            string  `json:"obj-parent"`
	IsVisible         bool    `json:"obj-visible"`
	IsEnabled         bool    `json:"obj-enabled"`
	Width             float32 `json:"obj-width"`
	Image             string  `json:"obj-image"`
	RelatedTable      string  `json:"related-table"`
	RelatedField      string  `json:"related-field"`
	InitValue         string  `json:"obj-initial-value"`
	UserDefinedType   string  `json:"user-def-type"`
	UserDefinedNumber int     `json:"user-def-num"`
	DropName          string  `json:"drop-name"`
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
	// "browse-cell-align": "",
	// "hotkey-code": "",
	// "hotkey-hdr-name": "",
	// "always-hidden": false,
	// "always-disabled": false,
	// "null-allowed": false,
	// "remove-line-feed": false,
	// "no-related-table": false,
	// "table-name": "",
	// "field-name": "",
	// "field-type": ""
}

func (s *ScreenDataModel) New() DataLoaderData {
	return &ScreenDataModel{}
}
