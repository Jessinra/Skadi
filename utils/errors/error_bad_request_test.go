package errors_test

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/errors"
)

const errMsgBadRequest = "it's a bad request error"

func Test_NewBadRequestError(t *testing.T) {
	t.Run("ok - created a bad request error", func(t *testing.T) {
		got := NewBadRequestError(errMsgBadRequest)

		assert.NotNil(t, got)
		assert.True(t, IsBadRequestError(got))
		assert.Equal(t, http.StatusBadRequest, got.Code())
		assert.EqualError(t, got, errMsgBadRequest)
	})

	t.Run("ok - created a bad request error with original error", func(t *testing.T) {
		err := errors.New("original error")
		got := NewBadRequestError(errMsgBadRequest, err)

		assert.NotNil(t, got)
		assert.True(t, IsBadRequestError(got))
		assert.Equal(t, http.StatusBadRequest, got.Code())
		assert.Equal(t, errMsgBadRequest, got.Message())
		assert.EqualError(t, got, "original error")
	})
}

func Test_IsBadRequestError(t *testing.T) {
	t.Run("ok - return true on NewBadRequestError", func(t *testing.T) {
		got := NewBadRequestError(errMsgBadRequest)
		assert.True(t, IsBadRequestError(got))
	})

	t.Run("ok - return true on custom error with code 400", func(t *testing.T) {
		got := NewCustomError(http.StatusBadRequest, errMsgBadRequest)
		assert.True(t, IsBadRequestError(got))
	})
}
