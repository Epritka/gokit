package validator

import (
	"encoding/json"
	"strconv"

	"github.com/Epritka/gokit/errors"
)

type Error struct {
	Message string
	Fields  []Field
	Debug   string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) AddFieldError(
	fieldName string, errorKey errors.ErrorKey,
) {
	e.Fields = append(
		e.Fields,
		NewFieldError(fieldName, errorKey),
	)
}

func (e *Error) AddFieldIndexError(
	fieldName string, errorKey errors.ErrorKey, index int,
) {
	e.Fields = append(
		e.Fields,
		NewFieldIndexError(fieldName, errorKey, index),
	)
}

func (e *Error) AddFieldOptionsError(
	fieldName string, errorKey errors.ErrorKey, options map[string]any,
) {
	e.Fields = append(
		e.Fields,
		NewFieldOptionsError(fieldName, errorKey, options),
	)
}

func (e *Error) AppendOptionsInFieldError(
	index int, options map[string]any,
) bool {
	if len(e.Fields)-1 > index {
		return false
	}

	if e.Fields[index].Options == nil {
		e.Fields[index].Options = options
		return true
	}

	for k, v := range options {
		e.Fields[index].Options[k] = v
	}
	return true
}

func (e *Error) AppendOptionsInLastFieldError(
	options map[string]any,
) bool {
	index := len(e.Fields) - 1
	return e.AppendOptionsInFieldError(index, options)
}

func (e *Error) SetOptionsInFieldError(
	index int, options map[string]any,
) bool {
	if len(e.Fields)-1 > index {
		return false
	}

	e.Fields[index].Options = options
	return true
}

func (e *Error) SetOptionsInLastFieldError(
	options map[string]any,
) bool {
	index := len(e.Fields) - 1
	return e.SetOptionsInFieldError(index, options)
}

func (e *Error) ClearOptionsInFieldError(
	index int, options map[string]any,
) bool {
	if len(e.Fields)-1 > index {
		return false
	}

	e.Fields[index].Options = nil
	return true
}

func (e *Error) ClearOptionsInLastFieldError(
	options map[string]any,
) bool {
	index := len(e.Fields) - 1
	return e.ClearOptionsInFieldError(index, options)
}

func (e *Error) SetIndexInFieldError(
	index, fieldIndex int,
) bool {
	if len(e.Fields)-1 > index {
		return false
	}

	e.Fields[index].Index = &fieldIndex
	return true
}

func (e *Error) SetIndexInLastFieldError(
	fieldIndex int,
) bool {
	index := len(e.Fields) - 1
	return e.SetIndexInFieldError(index, fieldIndex)
}

func (e *Error) ClearIndexInFieldError(
	index, fieldIndex int,
) bool {
	if len(e.Fields)-1 > index {
		return false
	}

	e.Fields[index].Options = nil
	return true
}

func (e *Error) ClearIndexInLastFieldError(
	fieldIndex int,
) bool {
	index := len(e.Fields) - 1
	return e.ClearIndexInFieldError(index, fieldIndex)
}

func (e *Error) SetMessageInFieldError(
	index int, message string,
) bool {
	if len(e.Fields)-1 > index {
		return false
	}

	e.Fields[index].Message = message
	return true
}

func (e *Error) SetMessageInLastFieldError(
	message string,
) bool {
	index := len(e.Fields) - 1
	return e.SetMessageInFieldError(index, message)
}

func (e *Error) MarshalJSON() ([]byte, error) {
	fields := map[string]map[string][]Field{}

	for _, v := range e.Fields {
		indexKey := "common"

		if v.Index != nil {
			indexKey = strconv.Itoa(*v.Index)
		}

		if _, ok := fields[indexKey]; !ok {
			fields[indexKey] = map[string][]Field{}
		}

		fields[indexKey][v.FieldName] = append(
			fields[indexKey][v.FieldName], v,
		)
	}

	if e.Message == "" {
		e.Message = errors.DefaultErorMessages[errors.ValidationErrorType]
	}

	if len(fields) == 1 {
		for _, v := range fields {
			return json.Marshal(&struct {
				ErrorType errors.ErrorType   `json:"type"`
				Message   string             `json:"message,omitempty"`
				Options   map[string]any     `json:"options,omitempty"`
				Fields    map[string][]Field `json:"fields,omitempty"`
			}{
				ErrorType: errors.ValidationErrorType,
				Message:   e.Message,
				Fields:    v,
			})
		}
	}

	return json.Marshal(&struct {
		ErrorType string                        `json:"type"`
		Message   string                        `json:"message,omitempty"`
		Options   map[string]any                `json:"options,omitempty"`
		Fields    map[string]map[string][]Field `json:"fields,omitempty"`
	}{
		Message: e.Message,
		Fields:  fields,
	})
}

// TODO Написать валидацию при конвертировании
func (e *Error) UnmarshalJSON(data []byte) error {
	type Error struct {
		Message string             `json:"message,omitempty"`
		Options map[string]any     `json:"options,omitempty"`
		Fields  map[string][]Field `json:"fields,omitempty"`
	}

	customError := Error{}
	if err := json.Unmarshal(data, &customError); err != nil {
		return err
	}

	e.Message = customError.Message
	e.Fields = []Field{}

	for k, fields := range customError.Fields {
		for _, f := range fields {
			e.Fields = append(e.Fields, Field{
				FieldName: k,
				ErrorKey:  f.ErrorKey,
				Message:   f.Message,
				Options:   f.Options,
			})
		}
	}

	return nil
}
