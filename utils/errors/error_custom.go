package errors

import "github.com/pkg/errors"

func NewCustomError(code int, message string, errs ...error) CustomError {
	return CustomError{
		code:    code,
		message: message,
		errs:    errs,
	}
}

func IsCustomError(err error) bool {
	return errors.As(err, &CustomError{})
}

type CustomError struct {
	code    int
	message string
	errs    []error
}

func (e CustomError) Code() int {
	return e.code
}

func (e CustomError) Message() string {
	return e.message
}

func (e CustomError) Error() string {
	if len(e.errs) > 0 {
		return e.errs[0].Error()
	}

	return e.Message()
}
