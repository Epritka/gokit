package errors

type (
	ErrorKey  string
	ErrorType string
)

const (
	UnknownErrorType ErrorType = "unknownError"

	UnauthorizedErrorType     ErrorType = "unauthorizedError"
	MethodNotAllowedErrorType ErrorType = "methodNotAllowedError"
	NotFoundErrorType         ErrorType = "notFoundError"
	ValidationErrorType       ErrorType = "validationError"

	InternalErrorType           ErrorType = "internalError"
	ServiceUnavailableErrorType ErrorType = "serviceUnavailable"
	DatabaseError               ErrorType = "databaseError"
)

var (
	DefaultErorMessages = map[ErrorType]string{
		UnknownErrorType:            "Unknown error",
		UnauthorizedErrorType:       "Authorization required",
		MethodNotAllowedErrorType:   "Method not allowed error",
		NotFoundErrorType:           "Not found error",
		ValidationErrorType:         "Validation error",
		InternalErrorType:           "Internal server error",
		ServiceUnavailableErrorType: "Service unavailable",
		DatabaseError:               "Database error",
	}
)
