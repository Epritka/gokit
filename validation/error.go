package validation

type Error struct {
	Fields []*Field `json:"fields"`
}

func (e *Error) Error() string {
	// TODO: correct string
	return "validation error"
}
