package entity

import (
	"strings"

	"github.com/Epritka/gokit/validation"
	"github.com/Epritka/gokit/validator"
)

type Role struct {
	Name string `json:"name"`
}

func (r *Role) Clear() {
	r.Name = strings.TrimSpace(r.Name)
}

func (r *Role) Fields() []*validator.Field {
	return []*validator.Field{
		validator.NewField("name", r.ValidateName),
	}
}

func (r *Role) ValidateName(field *validation.Field) error {
	if r.Name == "" {
		field.AddErrorKey(validation.Required)
	}
	return nil
}
