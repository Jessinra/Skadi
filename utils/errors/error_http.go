package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

type HTTPError interface {
	Code() int
	Message() string
	Error() string
}

func GetHTTPStatus(err error) int {
	err = errors.Cause(err)

	var httpErr HTTPError
	if ok := errors.As(err, &httpErr); ok {
		return httpErr.Code()
	}

	return http.StatusInternalServerError
}
