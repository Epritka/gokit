package validator

import "github.com/Epritka/gokit/errors"

type (
	Validator func(*ValidationError) (bool, error)
)

func CompareByFieldNameAndKey(err error, fieldName string, key errors.ErrorKey) bool {
	cErr, ok := err.(*ValidationError)
	if !ok {
		return false
	}

	for _, f := range cErr.Fields {
		if f.FieldName == fieldName && f.ErrorKey == key {
			return true
		}
	}

	return false
}

func Validate(clear func(), validators ...Validator) error {
	if clear != nil {
		clear()
	}

	validationError := ValidationError{
		Type: errors.ValidationErrorType,
	}

	for _, validator := range validators {
		isOver, err := validator(&validationError)
		if err != nil {
			return err
		}

		if isOver {
			return validationError.GetError()
		}
	}

	return validationError.GetError()
}
