package validation

type Error struct {
	Fields []*Field `json:"fields"`
}

func (e *Error) Error() string {
	return "validation error"
}
