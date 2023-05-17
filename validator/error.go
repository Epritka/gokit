package validator

import (
	"encoding/json"
	"strconv"

	"github.com/Epritka/gokit/errors"
)

type ValidationError struct {
	Message              string
	Type                 errors.ErrorType
	Fields               []Field
	IsErrorWithoutFields bool
	Debug                string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (e *ValidationError) GetError() error {
	if len(e.Fields) == 0 && !e.IsErrorWithoutFields {
		return nil
	}
	return e
}

func (e *ValidationError) AddFieldError(
	fieldName string, errorKey errors.ErrorKey,
) {
	e.Fields = append(
		e.Fields,
		NewFieldError(fieldName, errorKey),
	)
}

func (e *ValidationError) AddFieldIndexError(
	fieldName string, errorKey errors.ErrorKey, index int,
) {
	e.Fields = append(
		e.Fields,
		NewFieldIndexError(fieldName, errorKey, index),
	)
}

func (e *ValidationError) AddFieldOptionsError(
	fieldName string, errorKey errors.ErrorKey, options map[string]any,
) {
	e.Fields = append(
		e.Fields,
		NewFieldOptionsError(fieldName, errorKey, options),
	)
}

func (e *ValidationError) AppendOptionsInFieldError(
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

func (e *ValidationError) AppendOptionsInLastFieldError(
	options map[string]any,
) bool {
	index := len(e.Fields) - 1
	return e.AppendOptionsInFieldError(index, options)
}

func (e *ValidationError) SetOptionsInFieldError(
	index int, options map[string]any,
) bool {
	if len(e.Fields)-1 > index {
		return false
	}

	e.Fields[index].Options = options
	return true
}

func (e *ValidationError) SetOptionsInLastFieldError(
	options map[string]any,
) bool {
	index := len(e.Fields) - 1
	return e.SetOptionsInFieldError(index, options)
}

func (e *ValidationError) ClearOptionsInFieldError(
	index int, options map[string]any,
) bool {
	if len(e.Fields)-1 > index {
		return false
	}

	e.Fields[index].Options = nil
	return true
}

func (e *ValidationError) ClearOptionsInLastFieldError(
	options map[string]any,
) bool {
	index := len(e.Fields) - 1
	return e.ClearOptionsInFieldError(index, options)
}

func (e *ValidationError) SetIndexInFieldError(
	index, fieldIndex int,
) bool {
	if len(e.Fields)-1 > index {
		return false
	}

	e.Fields[index].Index = &fieldIndex
	return true
}

func (e *ValidationError) SetIndexInLastFieldError(
	fieldIndex int,
) bool {
	index := len(e.Fields) - 1
	return e.SetIndexInFieldError(index, fieldIndex)
}

func (e *ValidationError) ClearIndexInFieldError(
	index, fieldIndex int,
) bool {
	if len(e.Fields)-1 > index {
		return false
	}

	e.Fields[index].Options = nil
	return true
}

func (e *ValidationError) ClearIndexInLastFieldError(
	fieldIndex int,
) bool {
	index := len(e.Fields) - 1
	return e.ClearIndexInFieldError(index, fieldIndex)
}

func (e *ValidationError) SetMessageInFieldError(
	index int, message string,
) bool {
	if len(e.Fields)-1 > index {
		return false
	}

	e.Fields[index].Message = message
	return true
}

func (e *ValidationError) SetMessageInLastFieldError(
	message string,
) bool {
	index := len(e.Fields) - 1
	return e.SetMessageInFieldError(index, message)
}

func (e *ValidationError) MarshalJSON() ([]byte, error) {
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
		ErrorType: string(e.Type),
		Message:   e.Message,
		Fields:    fields,
	})
}

// TODO Написать валидацию при конвертировании
func (e *ValidationError) UnmarshalJSON(data []byte) error {
	type Error struct {
		Type    errors.ErrorType   `json:"type"`
		Message string             `json:"message,omitempty"`
		Options map[string]any     `json:"options,omitempty"`
		Fields  map[string][]Field `json:"fields,omitempty"`
	}

	customError := Error{}
	if err := json.Unmarshal(data, &customError); err != nil {
		return err
	}

	e.Type = customError.Type
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
