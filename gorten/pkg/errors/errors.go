package errors

import "net/http"

type CustomError struct {
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(statusCode int, message string) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func (e *CustomError) WithMessage(message string) *CustomError {
	return &CustomError{
		StatusCode: e.StatusCode,
		Message:    message,
	}
}

var (
	ErrUserNotFound            = NewCustomError(http.StatusNotFound, "user not found")
	ErrInternalServerError     = NewCustomError(http.StatusInternalServerError, "Internal Server Error")
	ErrContextDeadlineExceeded = NewCustomError(http.StatusRequestTimeout, "context deadline exceeded")
	ErrClientDisconnected      = NewCustomError(http.StatusGone, "client is disconnected")
	ErrFailedGenerateUUID      = NewCustomError(http.StatusInternalServerError, "failed to generate UUID")
	ErrSomethingWentWrong      = NewCustomError(http.StatusBadRequest, "something went wrong")
	ErrInvalidRequestPayload   = NewCustomError(http.StatusBadGateway, "invalid request payload")
	ErrValidationFailed        = NewCustomError(http.StatusBadRequest, "validation failed")
)
