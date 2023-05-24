package validation

type Field struct {
	Name   string
	Info   []Info
	Fields []*Field
	Index  *int
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
