package app

import (
	"fmt"
	"net/http"
)

type RespError struct {
	Code    int
	Message string
}

func NewRespError(code int, message string) *RespError {
	return &RespError{
		Code:    code,
		Message: message,
	}
}

func (e *RespError) Error() string {
	return e.Message
}

// BadRequest indicates client specified an invalid argument.
func BadRequest(format string, a ...interface{}) error {
	return NewRespError(http.StatusBadRequest, fmt.Sprintf(format, a...))
}

// NotFound means some requested model (e.g., file or directory) was
// not found.
func NotFound(format string, a ...interface{}) error {
	return NewRespError(http.StatusNotFound, fmt.Sprintf(format, a...))
}

// AlreadyExists means an attempt to create an model failed because one
// already exists.
func AlreadyExists(format string, a ...interface{}) error {
	return NewRespError(http.StatusConflict, fmt.Sprintf(format, a...))
}

// PermissionDenied indicates the caller does not have permission to
// execute the specified operation.
func PermissionDenied(format string, a ...interface{}) error {
	return NewRespError(http.StatusForbidden, fmt.Sprintf(format, a...))
}

// Aborted indicates the operation was aborted, typically due to a
// concurrency issue like sequencer check failures, 02transaction aborts, etc.
func Aborted(format string, a ...interface{}) error {
	return NewRespError(http.StatusConflict, fmt.Sprintf(format, a...))
}

// OutOfRange means operation was attempted past the valid range.
// E.g., seeking or reading past end of file.
func OutOfRange(format string, a ...interface{}) error {
	return NewRespError(http.StatusBadRequest, fmt.Sprintf(format, a...))
}

// Internal errors. Means some invariants expected by underlying
// system has been broken. If you see one of these errors,
// something is very broken.
func Internal(format string, a ...interface{}) error {
	return NewRespError(http.StatusInternalServerError, fmt.Sprintf(format, a...))
}

// DeadlineExceeded means operation expired before completion.
// For operations that change the state of the system, this error may be
// returned even if the operation has completed successfully. For
// example, a successful response from a server could have been delayed
// long enough for the deadline to expire.
func DeadlineExceeded(format string, a ...interface{}) error {
	return NewRespError(http.StatusGatewayTimeout, fmt.Sprintf(format, a...))
}

// Unimplemented indicates operation is not implemented or not
// supported/enabled in this service.
func Unimplemented(format string, a ...interface{}) error {
	return NewRespError(http.StatusNotImplemented, fmt.Sprintf(format, a...))
}

// ResourceExhausted indicates some resource has been exhausted, perhaps
// a per-user quota, or perhaps the entire file system is out of space.
func ResourceExhausted(format string, a ...interface{}) error {
	return NewRespError(http.StatusTooManyRequests, fmt.Sprintf(format, a...))
}

// FailedPrecondition indicates operation was rejected because the
// system is not in a state required for the operation's execution.
// For example, directory to be deleted may be non-empty, an rmdir
// operation is applied to a non-directory, etc.
func FailedPrecondition(format string, a ...interface{}) error {
	return NewRespError(http.StatusBadRequest, fmt.Sprintf(format, a...))
}
