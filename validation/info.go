package validation

import "github.com/Epritka/gokit/errors"

type Info struct {
	Key     errors.ErrorKey
	Options map[string]any
}
