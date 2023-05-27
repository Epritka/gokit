package user

import (
	"github.com/Epritka/gokit/example/internal/core/entity"
	"github.com/Epritka/gokit/example/internal/core/usecase/user/create"
	"github.com/Epritka/gokit/usecase"
)

type Interceptor struct {
}

func New() *Interceptor {
	return &Interceptor{}
}

func (i *Interceptor) Create(user *entity.User) error {
	return usecase.Run(create.New(user))
}
