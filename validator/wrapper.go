package validator

import (
	"encoding/json"
	"strconv"
)

func Wrap(err error) error {
	switch t := err.(type) {
	case *json.SyntaxError:
		return &Error{}
	case *json.UnmarshalTypeError:
		key := StandardTypes[t.Type.String()]
		if key == "" {
			key = NotType
		}

		return &Error{
			Fields: []Field{{
				FieldName: t.Field,
				ErrorKey:  key,
			}},
		}
	case *strconv.NumError:
		key := StrconvTypes[t.Func]
		if key == "" {
			key = NotType
		}

		return &Error{
			Fields: []Field{{
				ErrorKey: key,
			}},
		}
	default:
		return nil
	}
}
