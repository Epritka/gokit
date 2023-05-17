package validator

import (
	"encoding/json"
	"strconv"
)

func Wrap(err error) error {
	// TODO Дополнять ошибками
	switch t := err.(type) {
	case *json.SyntaxError:
		return &ValidationError{
			IsErrorWithoutFields: true,
		}
	case *json.UnmarshalTypeError:
		key := StandardTypes[t.Type.String()]
		if key == "" {
			key = NotType
		}

		return &ValidationError{
			Fields: []Field{{
				FieldName: t.Field,
				ErrorKey:  key,
			}},
		}
	case *strconv.NumError:
		// TODO Вынести куда нибудь, и подумать как лучше сделать эту ошибку
		// Наверное куда то в http ошибки
		return &ValidationError{
			Debug: t.Func,
		}
	default:
		return nil
	}
}
