package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewForbiddenError(message string, errs ...error) ForbiddenError {
	return ForbiddenError{
		message: message,
		errs:    errs,
	}
}

func IsForbiddenError(err error) bool {
	var e CustomError
	if errors.As(err, &e) && e.Code() == http.StatusForbidden {
		return true
	}

	return errors.As(err, &ForbiddenError{})
}

type ForbiddenError struct {
	message string
	errs    []error
}

func (ForbiddenError) Code() int {
	return http.StatusForbidden
}

func (e ForbiddenError) Message() string {
	return e.message
}

func (e ForbiddenError) Error() string {
	if len(e.errs) > 0 {
		return e.errs[0].Error()
	}

	return e.Message()
}
