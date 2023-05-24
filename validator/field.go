package validator

import (
	"encoding/json"

	"github.com/Epritka/gokit/errors"
)

type Field struct {
	ErrorKey  errors.ErrorKey
	FieldName string
	Message   string
	Options   map[string]any
	Index     *int // для привязки ошибок к элементам массива
}

func NewFieldError(fieldName string, errorKey errors.ErrorKey) Field {
	return Field{
		FieldName: fieldName,
		ErrorKey:  errorKey,
	}
}

func NewFieldIndexError(fieldName string, errorKey errors.ErrorKey, index int) Field {
	return Field{
		FieldName: fieldName,
		ErrorKey:  errorKey,
		Index:     &index,
	}
}

func NewFieldOptionsError(fieldName string, errorKey errors.ErrorKey, options map[string]any) Field {
	return Field{
		FieldName: fieldName,
		ErrorKey:  errorKey,
		Options:   options,
	}
}

func (f *Field) MarshalJSON() ([]byte, error) {
	if f.Message == "" {
		f.Message = DefaultFieldMessages[f.ErrorKey]
	}

	return json.Marshal(&struct {
		ErrorKey errors.ErrorKey `json:"key,omitempty"`
		Message  string          `json:"message,omitempty"`
		Options  map[string]any  `json:"options,omitempty"`
	}{
		ErrorKey: f.ErrorKey,
		Message:  f.Message,
		Options:  f.Options,
	})
}

func (f *Field) UnmarshalJSON(data []byte) error {
	type Field struct {
		ErrorKey errors.ErrorKey `json:"key,omitempty"`
		Message  string          `json:"message,omitempty"`
		Options  map[string]any  `json:"options,omitempty"`
	}

	fieldError := Field{}
	if err := json.Unmarshal(data, &fieldError); err != nil {
		return err
	}

	f.ErrorKey = fieldError.ErrorKey
	f.Message = fieldError.Message
	f.Options = fieldError.Options

	return nil
}
