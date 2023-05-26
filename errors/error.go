package errors

type DefaultError struct {
	Message     string         `json:"message,omitempty"`
	Type        ErrorType      `json:"type"`
	Meta        map[string]any `json:"meta,omitempty"`
	SourceError error          `json:"sourceError,omitempty"`
}

func (e *DefaultError) Error() string {
	if e.SourceError != nil {
		return e.SourceError.Error()
	}

	if e.Message != "" {
		return e.Message
	}

	return "it`s default error"
}
