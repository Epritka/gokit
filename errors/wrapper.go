package errors

import (
	"errors"
	"io"
	"net/url"
)

func NotFoundError() error {
	return &DefaultError{
		Type: NotFoundErrorType,
	}
}

func Wrap(err error) error {
	if errors.Is(err, io.EOF) {
		return &DefaultError{
			Type: ValidationErrorType,
		}
	}
	switch err.(type) {
	case *url.Error:
		return &DefaultError{
			Message:     DefaultErorMessages[ServiceUnavailableErrorType],
			Type:        ServiceUnavailableErrorType,
			SourceError: err,
		}
	default:
		return nil
	}
}
