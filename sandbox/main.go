package main

import (
	"encoding/json"
	"fmt"

	"github.com/Epritka/gokit/example"
	"github.com/Epritka/gokit/wrapper"
)

func getExampleUser() *example.User {
	return &example.User{
		Id:       -1,
		Name:     "name",
		Password: "Password",
		Doc:      nil,
		Roles: []*example.Role{
			{Name: "admin"},
			{Name: ""},
			{Name: "editor"},
		},
	}
}

func main() {
	err := example.Create(getExampleUser())
	if err != nil {
		code, response := wrapper.FailedHttpResponse(err)
		bytes, _ := json.Marshal(response)
		fmt.Println(code)
		fmt.Println(string(bytes))
		return
	}

	fmt.Println(200)
}
