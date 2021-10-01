package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

// NewUnauthorizedError create a new UnauthorizedError instance.
func NewUnauthorizedError(message string, errs ...error) UnauthorizedError {
	return UnauthorizedError{
		message: message,
		errs:    errs,
	}
}

func IsUnauthorizedError(err error) bool {
	var e CustomError
	if errors.As(err, &e) && e.Code() == http.StatusUnauthorized {
		return true
	}

	return errors.As(err, &UnauthorizedError{})
}

// UnauthorizedError will occurred if the data to save or modify have manadotry attribute thats not specified.
// Or not according some certain validity rule.
type UnauthorizedError struct {
	message string
	errs    []error
}

func (UnauthorizedError) Code() int {
	return http.StatusUnauthorized
}

func (e UnauthorizedError) Message() string {
	return e.message
}

func (e UnauthorizedError) Error() string {
	if len(e.errs) > 0 {
		return e.errs[0].Error()
	}

	return e.Message()
}
