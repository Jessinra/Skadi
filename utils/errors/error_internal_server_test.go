package errors_test

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/errors"
)

const errMsgInternalServer = "it's an internal server error"

func Test_NewInternalServerError(t *testing.T) {
	t.Run("ok - created an internal server error", func(t *testing.T) {
		got := NewInternalServerError(errMsgInternalServer)

		assert.NotNil(t, got)
		assert.True(t, IsInternalServerError(got))
		assert.Equal(t, http.StatusInternalServerError, got.Code())
		assert.EqualError(t, got, errMsgInternalServer)
	})

	t.Run("ok - created an internal server error with original error", func(t *testing.T) {
		err := errors.New("original error")
		got := NewInternalServerError(errMsgInternalServer, err)

		assert.NotNil(t, got)
		assert.True(t, IsInternalServerError(got))
		assert.Equal(t, http.StatusInternalServerError, got.Code())
		assert.Equal(t, errMsgInternalServer, got.Message())
		assert.EqualError(t, got, "original error")
	})
}

func Test_IsInternalServerError(t *testing.T) {
	t.Run("ok - return true on NewInternalServerError", func(t *testing.T) {
		got := NewInternalServerError(errMsgInternalServer)
		assert.True(t, IsInternalServerError(got))
	})

	t.Run("ok - return true on custom error with code 500", func(t *testing.T) {
		got := NewCustomError(http.StatusInternalServerError, errMsgInternalServer)
		assert.True(t, IsInternalServerError(got))
	})
}
