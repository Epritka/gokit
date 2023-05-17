package errors

type (
	ErrorKey  string
	ErrorType string
)

const (
	// [               General                  ]
	UnknownErrorType            ErrorType = "unknownError"
	NotFoundErrorType           ErrorType = "notFoundError"
	ServiceUnavailableErrorType ErrorType = "serviceUnavailable"

	// Internal
	InternalErrorType ErrorType = "internalError"

	DatabaseError = "databaseError"
	// ---------------------- //

	// HTTP
	HttpErrorType ErrorType = "httpError"

	UnauthorizedErrorType     ErrorType = "unauthorizedError"
	MethodNotAllowedErrorType ErrorType = "methodNotAllowedError"
	// ---------------------- //

	// Validation
	ValidationErrorType ErrorType = "validationError"
	// ---------------------- //
)

var (
	DefaultErorMessages = map[ErrorType]string{
		UnknownErrorType:            "Unknown error",
		NotFoundErrorType:           "Not found error",
		ServiceUnavailableErrorType: "Service unavailable",

		UnauthorizedErrorType:     "Authorization required",
		MethodNotAllowedErrorType: "Method not allowed error",

		ValidationErrorType: "Validation error",
	}
)
