package web

import (
	"github.com/pkg/errors"
	"net/http"
)

// FieldError is used to indicate an error with a specific request field.
type FieldError struct {
	Field string `json:"Field"`
	Error string `json:"Error"`
}

// ErrorResponse is the form used for API responses from failures in the API.
type ErrorResponse struct {
	Error  string       `json:"Error"`
	Fields []FieldError `json:"Fields,omitempty"`
}

// Error is used to pass an error during the request through the application
// with web specific context.
type Error struct {
	Err    error
	Status int
	Fields []FieldError
}

// NewError wraps the supplied error into a web.Error
// allowing mid.Errors to expose it to API
func NewError(err error) error {
	return &Error{err, http.StatusBadRequest, nil}
}

// NewRequestError wraps a provided error with an HTTP status code. This
// function should be used when handlers encounter expected errors.
func NewRequestError(err error, status int) error {
	return &Error{err, status, nil}
}

// Error implements the error interface. It uses the default message of the
// wrapped error. This is what will be shown in the services' logs.
func (err *Error) Error() string {
	return err.Err.Error()
}

// shutdown is a type used to help with the graceful termination of the service.
type shutdown struct {
	Message string
}

// NewShutdownError returns an error that causes the framework to signal a
// graceful shutdown.
func NewShutdownError(message string) error {
	return &shutdown{message}
}

// Error is the implementation of the error interface.
func (s *shutdown) Error() string {
	return s.Message
}

// IsShutdown checks to see if the shutdown error is contained in the specified
// error value.
func IsShutdown(err error) bool {
	if _, ok := errors.Cause(err).(*shutdown); ok {
		return true
	}
	return false
}
