package errors

import (
	"github.com/Epritka/gokit/localizer"
)

type (
	Cases map[string][]Case

	Case struct {
		ErrorKey  ErrorKey
		TextKey   localizer.TextKey
		Variables map[string]string
	}
)
