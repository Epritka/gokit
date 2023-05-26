package validator

import (
	"net"

	"github.com/Epritka/gokit/errors"
	"github.com/Epritka/gokit/validation"
	"golang.org/x/exp/constraints"
)

type (
	Ip   string
	Cidr string
	Asn  int
	Port int

	Number interface {
		constraints.Float | constraints.Integer
	}
)

func (cidr Cidr) Validate() (errors.ErrorKey, validation.Options) {
	ip, net, err := net.ParseCIDR(string(cidr))
	if err != nil {
		return validation.WrongFormat, nil
	}

	if net.IP.String() != ip.String() {
		return validation.NotMatch, validation.Options{
			"cidr": net.IP.String(),
		}
	}

	return "", nil
}

func (ip Ip) Validate() errors.ErrorKey {
	return ternary(net.ParseIP(string(ip)) == nil, validation.WrongFormat, "")
}

func (asn Asn) Validate() (errors.ErrorKey, validation.Options) {
	return MinMaxValidate(asn, 0, 65535)
}

func (port Port) Validate() (errors.ErrorKey, validation.Options) {
	return MinMaxValidate(port, 0, 65535)
}

func ternary[T any](cond bool, x T, y T) T {
	if cond {
		return x
	}
	return y
}

func MinMaxValidate[T Number](value, min, max T) (errors.ErrorKey, validation.Options) {
	if value < min {
		return validation.Min, validation.Options{"min": min}
	}

	if value > max {
		return validation.Max, validation.Options{"max": max}
	}

	return "", nil
}
