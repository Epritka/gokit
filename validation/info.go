package validation

import "github.com/Epritka/gokit/errors"

type Info struct {
	Key     errors.ErrorKey `json:"key"`
	Options map[string]any  `json:"options,omitempty"`
}
