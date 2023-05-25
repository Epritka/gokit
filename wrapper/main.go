package wrapper

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"strconv"

	kiterrors "github.com/Epritka/gokit/errors"
	"github.com/Epritka/gokit/validation"
)

type Wrapper func(error) error

func NotFoundError() error {
	return &kiterrors.DefaultError{
		Type: kiterrors.NotFoundErrorType,
	}
}

func Wrap(err error) error {
	if errors.Is(err, io.EOF) {
		return &kiterrors.DefaultError{
			Type: kiterrors.ValidationErrorType,
		}
	}

	switch t := err.(type) {
	case *validation.Error,
		*kiterrors.DefaultError:
		return err
	case *json.UnmarshalTypeError:
		key := validation.StandardTypes[t.Type.String()]
		if key == "" {
			key = validation.NotType
		}

		return &validation.Error{
			Fields: []*validation.Field{{
				Name: t.Field,
				Info: []validation.Info{{Key: key}},
			}}}

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
