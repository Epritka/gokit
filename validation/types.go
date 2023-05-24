package validation

import (
	"net"

	"github.com/Epritka/gokit/errors"
)

type (
	Ip   string
	Cidr string
	Asn  int
	Port int
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
)

var (
	DefaultFieldMessages = map[errors.ErrorKey]string{
		AlreadyExist: "object with field already exist",
		WrongFormat:  "field has wrong format",
		Required:     "field is required",

		MinLenght: "field too small",
		MaxLenght: "field too long",
		Min:       "field has value less minimal",
		Max:       "field has value more maximum",

		Email: "format of this field should be email",
	}
)

const (
	NotMatch    errors.ErrorKey = "notMatch"
	NotFound    errors.ErrorKey = "notFound"
	NotPossible errors.ErrorKey = "notPossible"

	NotType   errors.ErrorKey = "notType"
	NotInt    errors.ErrorKey = "notInt"
	NotFloat  errors.ErrorKey = "notFloat"
	NotBool   errors.ErrorKey = "notBool"
	NotArray  errors.ErrorKey = "notArray"
	NotString errors.ErrorKey = "notString"
)

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

func (cidr *Cidr) Validate() (errors.ErrorKey, map[string]any) {
	ip, net, err := net.ParseCIDR(string(*cidr))

	if err != nil {
		return WrongFormat, nil
	}

	if net.IP.String() != ip.String() {
		return NotMatch, map[string]any{
			"cidr": net.IP.String(),
		}
	}

	return "", nil
}

func (ip *Ip) Validate() errors.ErrorKey {
	if net.ParseIP(string(*ip)) != nil {
		return WrongFormat
	}
	return ""
}

func (asn *Asn) Validate() errors.ErrorKey {
	if *asn < 0 || *asn > 65535 {
		return WrongFormat
	}
	return ""
}

func (port *Port) Validate() errors.ErrorKey {
	if *port < 0 || *port > 65535 {
		return WrongFormat
	}
	return ""
}
