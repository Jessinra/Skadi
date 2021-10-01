package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewNotFoundError(message string, errs ...error) NotFoundError {
	return NotFoundError{
		message: message,
		errs:    errs,
	}
}

func IsNotFoundError(err error) bool {
	var e CustomError
	if errors.As(err, &e) && e.Code() == http.StatusNotFound {
		return true
	}

	return errors.As(err, &NotFoundError{})
}

// NotFoundError will occurred if theres a query operation into the system that yield no result.
type NotFoundError struct {
	message string
	errs    []error
}

func (NotFoundError) Code() int {
	return http.StatusNotFound
}

func (e NotFoundError) Message() string {
	return e.message
}

func (e NotFoundError) Error() string {
	if len(e.errs) > 0 {
		return e.errs[0].Error()
	}

	return e.Message()
}
