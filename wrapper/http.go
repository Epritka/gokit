package wrapper

import (
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
	return getStatusCodeByError(err), FailedResponse{Error: Wrap(err)}
}

func (r *FailedResponse) UnmarshalJSON(data []byte) error {
	validationError := struct{ Error validation.Error }{}
	err := json.Unmarshal(data, &validationError)
	if err == nil {
		r.Error = &validationError.Error
		return nil
	}

	deaultError := struct{ Error errors.DefaultError }{}
	err = json.Unmarshal(data, &deaultError)
	if err == nil {
		r.Error = &deaultError.Error
		return nil
	}

	return &errors.DefaultError{
		Message:     "error parse failed response",
		Type:        errors.InternalErrorType,
		SourceError: err,
	}
}

func (r *SuccessResponse[T]) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &r)
	if err != nil {
		return err
	}

	return nil
}

func getStatusCodeByError(err error) int {
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
