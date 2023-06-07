package wrapper

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"strconv"
	"strings"

	kiterrors "github.com/Epritka/gokit/errors"
	"github.com/Epritka/gokit/validation"
)

func NotFoundError() error {
	return &kiterrors.DefaultError{
		Type: kiterrors.NotFoundErrorType,
	}
}

func Wrap(err error) error {
	// TODO: Подумать куда переместить
	if errors.Is(err, io.EOF) {
		return &kiterrors.DefaultError{
			Type: kiterrors.ValidationErrorType,
		}
	}

	switch t := err.(type) {
	case *kiterrors.DefaultError,
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
	case *json.SyntaxError:
		return &kiterrors.DefaultError{
			Type:        kiterrors.ValidationErrorType,
			SourceError: err,
		}
	case *strconv.NumError:
		return &kiterrors.DefaultError{
			Message:     string(validation.StrconvTypes[t.Func]),
			Type:        kiterrors.ValidationErrorType,
			SourceError: err,
		}
	case *url.Error:
		return &kiterrors.DefaultError{
			Type:        kiterrors.ServiceUnavailableErrorType,
			SourceError: err,
		}
	default:
		return &kiterrors.DefaultError{
			Type:        kiterrors.UnknownErrorType,
			SourceError: err,
		}
	}
}
