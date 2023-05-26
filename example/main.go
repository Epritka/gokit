package example

import (
	"fmt"

	"github.com/Epritka/gokit/usecase"
	"github.com/Epritka/gokit/validation"
	"github.com/Epritka/gokit/validator"
)

type UseCase struct {
	dbUser UserInput
	User   UserInput
}

func (useCase *UseCase) validationId(field *validation.Field) error {
	// field.AddInfo(validation.Info{
	// 	Key: validation.Email,
	// })
	// return validator.Break

	// return errors.NotFoundError()
	// useCase.dbUser = UserInput{}
	return nil
}

func (useCase *UseCase) Fields() []*validator.Field {
	result := []*validator.Field{
		validator.NewField("id", useCase.validationId),
	}
	result = append(result, useCase.User.Fields()...)
	return result
}

func (useCase *UseCase) Execute() error {
	fmt.Println("exec")
	return nil
}

func Create() error {
	ip := "10.10.10.10"
	errorKey := validator.Ip(ip).Validate()
	fmt.Println(errorKey)

	useCase := &UseCase{
		User: UserInput{
			Name:      "name",
			Email:     "mail@mail.com",
			Prefix:    "10.10.10.100/32",
			IpAddress: ip,
			Password:  "Password",
			Role:      nil,
			Roles: []*RoleInput{
				{Name: "admin"},
				{Name: ""},
				{Name: "editor"},
			},
		},
	}

	return usecase.Run(useCase)
}
