package validation

type Error struct {
	Fields []*Field `json:"fields"`
}

func (e *Error) Error() string {
	return "validation error"
}

func FormattedFields(fields []*Field) map[string]any {
	formattedFields := map[string]any{}

	for _, field := range fields {
		name, fields := field.Formatted()
		formattedFields[name] = fields
	}

	return formattedFields
}
