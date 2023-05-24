package validator

import "github.com/Epritka/gokit/errors"

func CompareByFieldNameAndKey(err error, fieldName string, key errors.ErrorKey) bool {
	cErr, ok := err.(*Error)
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
