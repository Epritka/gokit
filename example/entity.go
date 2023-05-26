package example

import (
	"strings"

	"github.com/Epritka/gokit/validation"
	"github.com/Epritka/gokit/validator"
)

type User struct {
	Id       int
	Name     string
	Password string
	Roles    []*Role
	Doc      *Doc
}

type Role struct {
	Name string
}

type Doc struct {
	Name string
	Type string
}

func (u *User) Fields() []*validator.Field {
	return []*validator.Field{
		validator.NewField("name", u.ValidateName),
		validator.NewField("password", u.ValidatePassword),
		validator.NewSlice("roles", validator.SliceOfStruct(u.Roles)),
		validator.NewStruct("doc", u.Doc),
	}
}

func (u *User) ValidateName(field *validation.Field) error {
	if u.Name == "" {
		field.AddErrorKey(validation.Required)
	}
	return nil
}

func (u *User) ValidatePassword(field *validation.Field) error {
	if u.Password == "" {
		field.AddErrorKey(validation.Required)
	}
	return nil
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

func (d *Doc) Fields() []*validator.Field {
	return []*validator.Field{
		validator.NewField("name", d.ValidateName),
		validator.NewField("type", d.ValidateName),
	}
}

func (d *Doc) ValidateName(field *validation.Field) error {
	if d.Name == "" {
		field.AddErrorKey(validation.Required)
	}
	return nil
}

func (d *Doc) ValidateType(field *validation.Field) error {
	switch d.Type {
	case "passport",
		"driver's license":
	default:
		field.AddErrorKey(validation.UnknowType)
		return validator.Break
	}

	return nil
}
