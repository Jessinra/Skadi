package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

// NewInternalServerError creates new InternalServerError.
func NewInternalServerError(message string, errs ...error) InternalServerError {
	return InternalServerError{
		message: message,
		errs:    errs,
	}
}

func IsInternalServerError(err error) bool {
	var e CustomError
	if errors.As(err, &e) && e.Code() == http.StatusInternalServerError {
		return true
	}

	return errors.As(err, &InternalServerError{})
}

// InternalServerError is used to notify something wrong with the internal implementation.
type InternalServerError struct {
	message string
	errs    []error
}

func (InternalServerError) Code() int {
	return http.StatusInternalServerError
}

func (e InternalServerError) Message() string {
	return e.message
}

func (e InternalServerError) Error() string {
	if len(e.errs) > 0 {
		return e.errs[0].Error()
	}

	return e.Message()
}
