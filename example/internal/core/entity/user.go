package entity

import (
	"github.com/Epritka/gokit/validation"
	"github.com/Epritka/gokit/validator"
)

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Roles    []*Role `json:"roles"`
	Doc      *Doc    `json:"doc"`
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
