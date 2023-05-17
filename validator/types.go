package validator

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
	Required     errors.ErrorKey = "required"
	MaxString    errors.ErrorKey = "maxString"
	MinString    errors.ErrorKey = "minString"
	Min          errors.ErrorKey = "min"
	Email        errors.ErrorKey = "email"
	AlreadyExist errors.ErrorKey = "alreadyExist"
	WrongFormat  errors.ErrorKey = "wrongFormat"

	NotType     errors.ErrorKey = "notType"
	NotMatch    errors.ErrorKey = "notMatch"
	NotFound    errors.ErrorKey = "notFound"
	NotPossible errors.ErrorKey = "notPossible"
)

var (
	DefaultFieldMessages = map[errors.ErrorKey]string{
		Required:     "field is required",
		MaxString:    "field too long",
		MinString:    "field too small",
		Email:        "field has wrong format, need email",
		AlreadyExist: "object with field already exist",
		WrongFormat:  "field has wrong format",
	}
)

const (
	NotInt    errors.ErrorKey = "notInt"
	NotFloat  errors.ErrorKey = "notFloat"
	NotBool   errors.ErrorKey = "notBool"
	NotArray  errors.ErrorKey = "notArray"
	NotString errors.ErrorKey = "notString"
)

var (
	StandardTypes = map[string]errors.ErrorKey{
		"bool":    NotBool,
		"int":     NotInt,
		"int32":   NotInt,
		"int64":   NotInt,
		"uint":    NotInt,
		"uint32":  NotInt,
		"uint64":  NotInt,
		"float32": NotFloat,
		"float64": NotFloat,
		"slice":   NotArray,
		"string":  NotString,
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
