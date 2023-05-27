package main

import (
	"fmt"

	"github.com/Epritka/gokit/example/internal/core/entity"
	"github.com/Epritka/gokit/example/internal/core/usecase/user"
)

func getExampleUser() *entity.User {
	return &entity.User{
		Id:       -1,
		Name:     "name",
		Password: "Password",
		Doc:      nil,
		Roles: []*entity.Role{
			{Name: "admin"},
			{Name: ""},
			{Name: "editor"},
			{Name: "viewer"},
			{Name: ""},
		},
	}
}

func main() {
	err := user.New().Create(getExampleUser())
	fmt.Println(err)
}
