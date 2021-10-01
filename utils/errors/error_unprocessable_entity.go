package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewUnprocessableEntityError(message string, errs ...error) UnprocessableEntityError {
	return UnprocessableEntityError{
		message: message,
		errs:    errs,
	}
}

func IsUnprocessableEntityError(err error) bool {
	var e CustomError
	if errors.As(err, &e) && e.Code() == http.StatusUnprocessableEntity {
		return true
	}

	return errors.As(err, &UnprocessableEntityError{})
}

type UnprocessableEntityError struct {
	message string
	errs    []error
}

func (UnprocessableEntityError) Code() int {
	return http.StatusUnprocessableEntity
}

func (e UnprocessableEntityError) Message() string {
	return e.message
}

func (e UnprocessableEntityError) Error() string {
	if len(e.errs) > 0 {
		return e.errs[0].Error()
	}

	return e.Message()
}
