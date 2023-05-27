package create

import (
	"fmt"
)

func (useCase *UseCase) Execute() error {
	fmt.Println("executed")
	return nil
}
