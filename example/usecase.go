package example

type UseCase struct {
	User User
}

func NewUseCase(user *User) *UseCase {
	return &UseCase{User: *user}
}
