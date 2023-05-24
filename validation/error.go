package validation

type Error struct {
	Fields []*Field
}

func (e *Error) Error() string {
	return "validation error"
}
