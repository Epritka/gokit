package entity

import (
	"github.com/Epritka/gokit/validation"
	"github.com/Epritka/gokit/validator"
)

type Doc struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Test Test   `json:"test"`
}

type Test struct {
	Test Test2 `json:"test"`
}

type Test2 struct {
	Name string `json:"name"`
}

func (d *Doc) Fields() validator.Fields {
	return validator.NewFields(
		validator.NewField("name", d.ValidateName),
		validator.NewField("type", d.ValidateName),
	)
}

func (d *Doc) ValidateName(field *validation.Field) error {
	if d.Name == "" {
		field.AddErrorKey(validation.Required)
	}
	return nil
}

func (d *Doc) ValidateType(field *validation.Field) error {
	switch d.Type {
	case "passport", "driver's license":
	default:
		field.AddErrorKey(validation.NotEnum)
		return validator.Break
	}

	return nil
}
