package wrapper

import (
	"github.com/Epritka/gokit/errors"
)

type Wrapper func(error) error

func Wrap(err error, wrappers ...Wrapper) error {
	result := errors.Wrap(err)
	if result != nil {
		return result
	}

	// result = validator.Wrap(err)
	// if result != nil {
	// 	return result
	// }

	for _, wrapper := range wrappers {
		result = wrapper(err)
		if result != nil {
			return result
		}
	}

	return &errors.DefaultError{
		SourceError: err,
		Type:        errors.UnknownErrorType,
	}
}
