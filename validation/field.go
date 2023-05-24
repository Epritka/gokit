package validation

type Field struct {
	Name   string   `json:"name"`
	Info   []Info   `json:"info"`
	Fields []*Field `json:"fields,omitempty"`
	Index  *int     `json:"index,omitempty"`
}

func NewField(name string) *Field {
	return &Field{
		Name:   name,
		Info:   []Info{},
		Fields: []*Field{},
	}
}

func (item *Field) IsEmpty() bool {
	return len(item.Fields) == 0 && len(item.Info) == 0
}

func (item *Field) AddInfo(info Info) {
	item.Info = append(item.Info, info)
}
