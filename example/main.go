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
	field.AddInfo(validation.Info{
		Key: validation.Email,
	})

	// return errors.NotFoundError()
	useCase.dbUser = UserInput{}
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
	useCase := &UseCase{
		User: UserInput{
			Name:      "name",
			Email:     "mail@mail.com",
			IpAddress: "10.10.10.10",
			Password:  "Password",
			Roles: []*RoleInput{
				{Name: "admin"},
				{Name: ""},
				{Name: "editor"},
			},
		},
	}

	return usecase.Run(useCase)
}
