package errors

import (
	stderrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

func New(message string) error {
	return stderrors.New(message) //nolint do not define dynamic errors, use wrapped static errors instead: "errors.New(message)"
}

func Newf(format string, a ...interface{}) error {
	return stderrors.New(fmt.Sprintf(format, a...)) //nolint do not define dynamic errors, use wrapped static errors instead: "errors.New(message)"
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

func Cause(err error) error {
	return errors.Cause(err)
}
