package validation

import (
	"github.com/Epritka/gokit/errors"
)

type (
	Options map[string]any
	Info    struct {
		Key     errors.ErrorKey `json:"key"`
		Options Options         `json:"options,omitempty"`
	}

	Field struct {
		Name        string        `json:"name"`
		Info        []Info        `json:"info,omitempty"`
		Fields      []*Field      `json:"fields,omitempty"`
		ArrayFields []*ArrayField `json:"arrayFields,omitempty"`
	}

	ArrayField struct {
		Fields []*Field `json:"fields"`
		Index  int      `json:"index"`
	}
)

func NewField(name string) *Field {
	return &Field{
		Name:        name,
		Info:        []Info{},
		Fields:      []*Field{},
		ArrayFields: []*ArrayField{},
	}
}

func NewArrayField(fields []*Field, index int) *ArrayField {
	return &ArrayField{
		Fields: fields,
		Index:  index,
	}
}

func (field *Field) IsEmpty() bool {
	return len(field.Info) == 0 &&
		len(field.Fields) == 0 &&
		len(field.ArrayFields) == 0
}

func (field *Field) AppendField(f *Field) {
	field.Fields = append(field.Fields, f)
}

func (field *Field) AddInfo(info Info) {
	field.Info = append(field.Info, info)
}

func (field *Field) AddErrorKey(key errors.ErrorKey) {
	field.Info = append(field.Info, Info{Key: key})
}

func (field *Field) AddErrorKeyOptions(key errors.ErrorKey, options Options) {
	field.Info = append(field.Info, Info{Key: key, Options: options})
}
