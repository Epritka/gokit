package example

import (
	"github.com/Epritka/gokit/validation"
	"github.com/Epritka/gokit/validator"
)

type UserInput struct {
	Name      string
	IpAddress string
	Email     string
	Password  string
	Roles     []*RoleInput
}

func (u *UserInput) Fields() []*validator.Field {
	return []*validator.Field{
		validator.NewField("ipAddress", u.ValidateIpAddress),
		validator.NewSlice("roles", validator.SliceOfStruct(u.Roles)),
	}
}

func (u *UserInput) ValidateIpAddress(field *validation.Field) error {
	if u.IpAddress == "" {
		field.AddInfo(validation.Info{
			Key: validation.Required,
		})
		return nil
	}

	ip := validation.Ip(u.IpAddress)
	errorKey := ip.Validate()

	if errorKey != "" {
		field.AddInfo(validation.Info{
			Key: validation.WrongFormat,
		})
	}

	return nil
}
