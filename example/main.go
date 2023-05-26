package example

import "github.com/Epritka/gokit/usecase"

func Create(user *User) error {
	return usecase.Run(NewUseCase(user))
}
