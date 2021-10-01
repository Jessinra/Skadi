package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

func NewResourceAlreadyExistsError(message string, errs ...error) ResourceAlreadyExistsError {
	return ResourceAlreadyExistsError{
		message: message,
		errs:    errs,
	}
}

func IsResourceAlreadyExistsError(err error) bool {
	var e CustomError
	if errors.As(err, &e) && e.Code() == http.StatusConflict {
		return true
	}

	return errors.As(err, &ResourceAlreadyExistsError{})
}

// ResourceAlreadyExistsError will occur if there's a query wants to create a new resource,
// but the same resource already exists.
type ResourceAlreadyExistsError struct {
	message string
	errs    []error
}

func (ResourceAlreadyExistsError) Code() int {
	return http.StatusConflict
}

func (e ResourceAlreadyExistsError) Message() string {
	return e.message
}

func (e ResourceAlreadyExistsError) Error() string {
	if len(e.errs) > 0 {
		return e.errs[0].Error()
	}

	return e.Message()
}
