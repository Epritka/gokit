package main

import (
	"fmt"

	"github.com/Epritka/gokit/test/validator"
	"github.com/Epritka/gokit/usecase"
	"github.com/Epritka/gokit/validation"
)

type UseCase struct {
	dbUser UserInput
	User   UserInput
}

func (useCase *UseCase) validationId(field *validation.Field) error {
	// user, err := fields.userRepository.Get(fields.Id)
	// if err != nil {
	// 	vErr.Type = coreError.NotFoundErrorType
	// 	vErr.IsErrorWithoutFields = true
	// 	return true
	// }
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

func main() {
	useCase := UseCase{
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

	useCase.Execute()
	err := usecase.Run(useCase)
	fmt.Println(err.Error())
}
