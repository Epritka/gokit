package main

import (
	"strings"

	"github.com/Epritka/gokit/test/validator"
	"github.com/Epritka/gokit/validation"
)

type RoleInput struct {
	Name string
}

func (r *RoleInput) Clear() {
	r.Name = strings.TrimSpace(r.Name)
}

func (r *RoleInput) Fields() []*validator.Field {
	return []*validator.Field{
		validator.NewField("name", r.ValidateName),
	}
}

func (r *RoleInput) ValidateName(field *validation.Field) error {
	if r.Name == "" {
		field.AddInfo(validation.Info{
			Key: validation.Required,
		})
		return nil
	}

	return nil
}
