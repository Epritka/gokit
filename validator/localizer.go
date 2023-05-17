package validator

import (
	"fmt"

	"github.com/Epritka/gokit/errors"
	"github.com/Epritka/gokit/localizer"
)

func (e *ValidationError) Localize(
	localizer localizer.Localizer,
	localizeMessage func(string, errors.ErrorType) string,
	cases []errors.Cases,
) error {
	fields := []Field{}
	for _, f := range e.Fields {
		message := ""
		for _, _case := range cases {
			if findedCases, find := _case[f.FieldName]; find {
				for _, c := range findedCases {
					if c.ErrorKey != f.ErrorKey {
						continue
					}
					if len(f.Options) > 0 {
						message = localizer.GetTextByKeyWithVariables(c.TextKey, overloadOptions(f.Options, c.Variables))
					} else {
						message = localizer.GetTextByKey(c.TextKey)
					}
				}
			}
		}

		// TODO потом заменить, когда все работать будет (нужно для дебага)
		if message == "" || message == localizer.GetDefaultText() {
			message = fmt.Sprintf("DEBUG:  message: (%s), errorKey: (%s), options: (%v)", f.Message, f.ErrorKey, f.Options)
			// message = localizer.GetDefaultText()
		}

		fields = append(fields, Field{
			FieldName: f.FieldName,
			ErrorKey:  f.ErrorKey,
			Message:   message,
			Options:   f.Options,
			Index:     f.Index,
		})
	}

	return &ValidationError{
		Message: localizeMessage(e.Message, e.Type),
		// TODO потом заменить, когда все работать будет (нужно для дебага)
		Debug:                e.Message,
		Type:                 e.Type,
		Fields:               fields,
		IsErrorWithoutFields: e.IsErrorWithoutFields,
	}
}

func overloadOptions(options map[string]any, variables map[string]string) map[string]any {
	newOptions := map[string]any{}
	for k, v := range options {
		if newKey, find := variables[k]; find {
			k = newKey
		}
		newOptions[k] = v
	}

	return newOptions
}
