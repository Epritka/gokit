package errors

import "encoding/json"

type DefaultError struct {
	Message     string
	Type        ErrorType
	SourceError error
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
	if e.SourceError != nil {
		e.Message = e.SourceError.Error()
	}

	if e.Message == "" {
		e.Message = DefaultErorMessages[e.Type]
	}

	return json.Marshal(&struct {
		Message     string    `json:"message,omitempty"`
		Type        ErrorType `json:"type"`
		SourceError error     `json:"sourceError,omitempty"`
	}{
		Message:     e.Message,
		Type:        e.Type,
		SourceError: e.SourceError,
	})
}
