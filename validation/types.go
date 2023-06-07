package validation

import (
	"github.com/Epritka/gokit/errors"
)

const (
	AlreadyExist errors.ErrorKey = "alreadyExist"
	WrongFormat  errors.ErrorKey = "wrongFormat"
	Required     errors.ErrorKey = "required"

	MinLenght errors.ErrorKey = "minLenght"
	MaxLenght errors.ErrorKey = "maxLenght"
	Min       errors.ErrorKey = "min"
	Max       errors.ErrorKey = "max"

	Email errors.ErrorKey = "email"

	NotMatch    errors.ErrorKey = "notMatch"
	NotFound    errors.ErrorKey = "notFound"
	NotPossible errors.ErrorKey = "notPossible"

	NotEnum   errors.ErrorKey = "notEnum"
	NotType   errors.ErrorKey = "notType"
	NotInt    errors.ErrorKey = "notInt"
	NotFloat  errors.ErrorKey = "notFloat"
	NotBool   errors.ErrorKey = "notBool"
	NotArray  errors.ErrorKey = "notArray"
	NotString errors.ErrorKey = "notString"
)

var DefaultFieldMessages = map[errors.ErrorKey]string{
	AlreadyExist: "object with field already exist",
	WrongFormat:  "field has wrong format",
	Required:     "field is required",

	MinLenght: "field too small",
	MaxLenght: "field too long",
	Min:       "field has value less minimal",
	Max:       "field has value more maximum",

	Email: "format of this field should be email",
}

var (
	StandardTypes = map[string]errors.ErrorKey{
		"int":     NotInt,
		"int32":   NotInt,
		"int64":   NotInt,
		"uint":    NotInt,
		"uint32":  NotInt,
		"uint64":  NotInt,
		"float32": NotFloat,
		"float64": NotFloat,
		"bool":    NotBool,
		"slice":   NotArray,
		"string":  NotString,
	}

	StrconvTypes = map[string]errors.ErrorKey{
		"ParseComplex": NotType,
		"ParseInt":     NotInt,
		"ParseUint":    NotInt,
		"ParseFloat":   NotFloat,
		"ParseBool":    NotBool,
	}
)
