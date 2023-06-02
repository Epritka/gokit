package wrapper

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Epritka/gokit/errors"
	"github.com/Epritka/gokit/validation"
)

type SuccessResponse[T any] struct {
	Data  T   `json:"data"`
	Count int `json:"count,omitempty"`
}

type FailedResponse struct {
	Error error `json:"error"`
}

func SuccessHttpResponse[T any](data T) (int, SuccessResponse[T]) {
	return http.StatusOK, SuccessResponse[T]{Data: data}
}

func SuccessHttpResponseWithCount[T any](data T, count int) (int, SuccessResponse[T]) {
	return http.StatusOK, SuccessResponse[T]{Data: data, Count: count}
}

func FailedHttpResponse(err error) (int, FailedResponse) {
	err = Wrap(err)
	return GetStatusCodeByError(err), FailedResponse{Error: err}
}

func ErrorFromFaildedResponse(data []byte) error {
	response := FailedResponse{}
	err := json.Unmarshal(data, &response)
	if err != nil {
		return err
	}
	return response.Error
}

func (r *FailedResponse) UnmarshalJSON(data []byte) error {
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	validationError := struct{ Error validation.Error }{}

	err := decoder.Decode(&validationError)
	if err == nil {
		r.Error = &validationError.Error
		return nil
	}

	deaultError := struct{ Error errors.DefaultError }{}
	err = decoder.Decode(&deaultError)
	if err == nil {
		r.Error = &deaultError.Error
		return nil
	}

	meta := map[string]any{}
	decoder.Decode(&meta)

	return &errors.DefaultError{
		Type: errors.InternalErrorType,
		Meta: meta,
	}
}

func GetStatusCodeByError(err error) int {
	errType := errors.UnknownErrorType

	switch t := err.(type) {
	case *errors.DefaultError:
		errType = t.Type
	case *validation.Error:
		return http.StatusUnprocessableEntity
	}

	switch errType {
	case errors.UnauthorizedErrorType:
		return http.StatusUnauthorized

	case errors.NotFoundErrorType:
		return http.StatusNotFound

	case errors.ValidationErrorType:
		return http.StatusUnprocessableEntity

	case errors.MethodNotAllowedErrorType:
		return http.StatusMethodNotAllowed

	case errors.ServiceUnavailableErrorType,
		errors.UnknownErrorType:
		fallthrough
	default:
		return http.StatusInternalServerError
	}
}
