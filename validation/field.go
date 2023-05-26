package validation

import "github.com/Epritka/gokit/errors"

type Field struct {
	Name   string   `json:"name"`
	Info   []Info   `json:"info,omitempty"`
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

func (item *Field) AddErrorKey(key errors.ErrorKey) {
	item.Info = append(item.Info, Info{Key: key})
}

func (item *Field) AddErrorKeyOptions(key errors.ErrorKey, options Options) {
	item.Info = append(item.Info, Info{Key: key, Options: options})
}
