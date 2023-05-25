package validator

import (
	"fmt"

	"github.com/Epritka/gokit/validation"
)

type (
	ValidateFunc func(*validation.Field) error
)

var (
	Break = fmt.Errorf("break")
)

func Validate(structure Structure) error {
	fields, err := validate(structure)

	if err != nil {
		return err
	}

	if len(fields) > 0 {
		return &validation.Error{
			Fields: fields,
		}
	}

	return nil
}

func validate(structure Structure) ([]*validation.Field, error) {
	fields := []*validation.Field{}
	if structure == nil {
		return fields, nil
	}

	if clearer, ok := structure.(interface{ Clear() }); ok {
		clearer.Clear()
	}

	for _, f := range structure.Fields() {
		field := validation.NewField(f.name)

		switch f.fieldType {
		case primitiveType:
			err := f.validateFunc(field)
			if err != nil {
				if err.Error() != "break" {
					return nil, err
				}

				if !field.IsEmpty() {
					fields = append(fields, field)
				}

				return fields, nil
			}
		case structureType:
			fs, err := validate(f.structure)
			if err != nil {
				return nil, err
			}

			field.Fields = append(field.Fields, fs...)
		case sliceType:
			for i, structure := range f.slice {
				fs, err := validate(structure)
				if err != nil {
					return nil, err
				}

				if !field.IsEmpty() {
					field.Index = &i
				}

				field.Fields = append(field.Fields, fs...)
			}
		}

		if !field.IsEmpty() {
			fields = append(fields, field)
		}
	}

	return fields, nil
}
