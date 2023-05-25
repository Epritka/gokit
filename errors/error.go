package errors

import "encoding/json"

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

func (e *DefaultError) MarshalJSON() ([]byte, error) {
	if e.Message == "" {
		e.Message = DefaultErorMessages[e.Type]
	}

	return json.Marshal(e)
}
