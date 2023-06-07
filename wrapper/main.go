package wrapper

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/Epritka/gokit/errors"
	"github.com/Epritka/gokit/validation"
)

func NotFoundError() error {
	return &errors.DefaultError{
		Type: errors.NotFoundErrorType,
	}
}

func HttpWrap(err error) error {
	if err == io.EOF {
		return &errors.DefaultError{
			Type: errors.ValidationErrorType,
		}
	}

	switch t := err.(type) {
	case *json.SyntaxError:
		return &errors.DefaultError{
			Type:        errors.ValidationErrorType,
			SourceError: err,
		}
	case *strconv.NumError:
		return &errors.DefaultError{
			Message:     string(validation.StrconvTypes[t.Func]),
			Type:        errors.ValidationErrorType,
			SourceError: err,
		}
	}

	return Wrap(err)
}

func Wrap(err error) error {
	switch t := err.(type) {
	case *errors.DefaultError,
		*validation.Error:
		return err
	case *json.UnmarshalTypeError:
		key := validation.StandardTypes[t.Type.String()]
		if key == "" {
			key = validation.NotType
		}

		fieldsNames := strings.Split(t.Field, ".")
		mainField := validation.NewField(fieldsNames[0])
		currentField := mainField
		size := len(fieldsNames)

		for i := 1; i < size; i++ {
			field := validation.NewField(fieldsNames[i])
			currentField.AppendField(field)
			currentField = field
		}

		currentField.AddErrorKey(key)

		return &validation.Error{
			Fields: []*validation.Field{mainField},
		}
	case *url.Error:
		return &errors.DefaultError{
			Type:        errors.ServiceUnavailableErrorType,
			SourceError: err,
		}
	default:
		return &errors.DefaultError{
			Type:        errors.UnknownErrorType,
			SourceError: err,
		}
	}
}
