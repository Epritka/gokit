package create

import (
	"github.com/Epritka/gokit/validation"
	"github.com/Epritka/gokit/validator"
)

func (useCase *UseCase) Fields() validator.Fields {
	return validator.NewFields(
		validator.NewField("id", useCase.validationId),
	).Join(&useCase.User)
}

func (useCase *UseCase) validationId(field *validation.Field) error {
	if useCase.User.Id < 0 {
		field.AddErrorKey(validation.Min)
	}

	return nil
}
