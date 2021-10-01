package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewNotImplementedError(message string, errs ...error) NotImplementedError {
	return NotImplementedError{
		message: message,
		errs:    errs,
	}
}

func IsNotImplementedError(err error) bool {
	var e CustomError
	if errors.As(err, &e) && e.Code() == http.StatusNotImplemented {
		return true
	}

	return errors.As(err, &NotImplementedError{})
}

type NotImplementedError struct {
	message string
	errs    []error
}

func (NotImplementedError) Code() int {
	return http.StatusNotImplemented
}

func (e NotImplementedError) Message() string {
	return e.message
}

func (e NotImplementedError) Error() string {
	if len(e.errs) > 0 {
		return e.errs[0].Error()
	}

	return e.Message()
}
