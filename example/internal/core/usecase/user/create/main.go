package create

import "github.com/Epritka/gokit/example/internal/core/entity"

type UseCase struct {
	User entity.User
}

func New(user *entity.User) *UseCase {
	return &UseCase{User: *user}
}
