package usecase

import "github.com/Epritka/gokit/test/validator"

func Run(useCase any) (err error) {
	if initializator, ok := useCase.(interface{ Init() error }); ok {
		return initializator.Init()
	}

	if structure, ok := useCase.(validator.Structure); ok {
		if err := validator.Validate(structure); err != nil {
			return err
		}
	}

	if executer, ok := useCase.(interface{ Execute() error }); ok {
		return executer.Execute()
	}

	return
}

func RunOutput[Output any](useCase any) (output Output, err error) {
	if initializator, ok := useCase.(interface{ Init() error }); ok {
		return output, initializator.Init()
	}

	if structure, ok := useCase.(validator.Structure); ok {
		if err := validator.Validate(structure); err != nil {
			return output, err
		}
	}

	if executer, ok := useCase.(interface{ Execute() (Output, error) }); ok {
		return executer.Execute()
	}

	return
}
