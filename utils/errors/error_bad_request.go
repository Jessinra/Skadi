package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

// NewBadRequestError create a new BadRequestError instance.
func NewBadRequestError(message string, errs ...error) BadRequestError {
	return BadRequestError{
		message: message,
		errs:    errs,
	}
}

func IsBadRequestError(err error) bool {
	var e CustomError
	if errors.As(err, &e) && e.Code() == http.StatusBadRequest {
		return true
	}

	return errors.As(err, &BadRequestError{})
}

// BadRequestError will occurred if the data to save or modify have manadotry attribute thats not specified.
// Or not according some certain validity rule.
type BadRequestError struct {
	message string
	errs    []error
}

func (BadRequestError) Code() int {
	return http.StatusBadRequest
}

func (e BadRequestError) Message() string {
	return e.message
}

func (e BadRequestError) Error() string {
	if len(e.errs) > 0 {
		return e.errs[0].Error()
	}

	return e.Message()
}
