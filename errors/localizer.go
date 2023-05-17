package errors

import (
	"github.com/Epritka/gokit/localizer"
)

// func Localize(
// 	localizer localizer.Localizer,
// 	localizeMessage func(ErrorType) string,
// 	cases []Cases,
// ) error

func (e *DefaultError) Localize(
	localizer localizer.Localizer,
	localizeMessage func(ErrorType) string,
	cases []Cases,
) error {
	return &DefaultError{
		Message:     localizeMessage(e.Type),
		SourceError: e.SourceError,
		Type:        e.Type,
	}
}
