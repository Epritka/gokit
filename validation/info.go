package validation

import "github.com/Epritka/gokit/errors"

type Options map[string]any
type Info struct {
	Key     errors.ErrorKey `json:"key"`
	Options Options         `json:"options,omitempty"`
}
