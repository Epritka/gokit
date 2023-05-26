package example

import (
	"fmt"
	"net"

	"github.com/Epritka/gokit/validation"
	"github.com/Epritka/gokit/validator"
)

type UserInput struct {
	Name      string
	Prefix    string
	IpAddress string
	Email     string
	Password  string
	Roles     []*RoleInput
	Role      *RoleInput
}

func (u *UserInput) Fields() []*validator.Field {
	return []*validator.Field{
		validator.NewField("prefix", u.ValidatePrefix),
		validator.NewField("ipAddress", u.ValidateIpAddress),
		validator.NewStruct("role", u.Role),
		validator.NewSlice("roles", validator.SliceOfStruct(u.Roles)),
	}
}

func (u *UserInput) ValidatePrefix(field *validation.Field) error {
	errorKey, options := validator.Cidr(u.Prefix).Validate()
	if errorKey != "" {
		field.AddErrorKeyOptions(errorKey, options)
		return validator.Break
	}
	return nil
}

func (u *UserInput) ValidateIpAddress(field *validation.Field) error {
	if u.IpAddress == "" {
		field.AddInfo(validation.Info{
			Key: validation.Required,
		})
		return nil
	}

	errorKey := validator.Ip(u.IpAddress).Validate()
	if errorKey != "" {
		fmt.Println(u.IpAddress)
		field.AddErrorKey(errorKey)
	}
	_, ipNet, _ := net.ParseCIDR(u.Prefix)
	if !ipNet.Contains(net.ParseIP(u.IpAddress)) {
		field.AddErrorKey(validation.NotMatch)
	}

	return nil
}
