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

func (cidr Cidr) Validate() (errors.ErrorKey, map[string]any) {
	ip, net, err := net.ParseCIDR(string(cidr))
	if err != nil {
		return validation.WrongFormat, nil
	}

	if net.IP.String() != ip.String() {
		return validation.NotMatch, map[string]any{
			"cidr": net.IP.String(),
		}
	}

	return "", nil
}

func (ip Ip) Validate() errors.ErrorKey {
	return ternary(net.ParseIP(string(ip)) == nil, validation.WrongFormat, "")
}

func (asn Asn) Validate() errors.ErrorKey {
	return MinMaxValidate(asn, 0, 65535)
}

func (port Port) Validate() errors.ErrorKey {
	return MinMaxValidate(port, 0, 65535)
}

func ternary[T any](cond bool, x T, y T) T {
	if cond {
		return x
	}
	return y
}

func MinMaxValidate[T Number](value, min, max T) errors.ErrorKey {
	if value < min {
		return validation.Min
	}

	if value > max {
		return validation.Max
	}

	return ""
}
