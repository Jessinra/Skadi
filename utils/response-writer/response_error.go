package writer

import (
	"net/http"
	"time"

	"gitlab.com/trivery-id/skadi/utils/errors"
)

var ErrMsgUnableToBindJSON = "unable to bind JSON"

type ErrorResponse struct {
	UUID      string    `json:"uuid,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`

	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewErrorResponse(err error) ErrorResponse {
	var httpErr errors.HTTPError
	if errors.As(err, &httpErr) {
		return ErrorResponse{
			Message: httpErr.Message(),
			Status:  httpErr.Code(),
			Error:   httpErr.Error(),
		}
	}

	return ErrorResponse{
		Message: "Internal server error occurred",
		Status:  http.StatusInternalServerError,
		Error:   err.Error(),
	}
}
