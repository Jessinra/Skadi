package errors_test

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/errors"
)

const errMsgUnprocessableEntity = "it's an unprocessable entity error"

func Test_NewUnprocessableEntityError(t *testing.T) {
	t.Run("ok - created an unprocessable entity error", func(t *testing.T) {
		got := NewUnprocessableEntityError(errMsgUnprocessableEntity)

		assert.NotNil(t, got)
		assert.True(t, IsUnprocessableEntityError(got))
		assert.Equal(t, http.StatusUnprocessableEntity, got.Code())
		assert.EqualError(t, got, errMsgUnprocessableEntity)
	})

	t.Run("ok - created an unprocessable entity error with original error", func(t *testing.T) {
		err := errors.New("original error")
		got := NewUnprocessableEntityError(errMsgUnprocessableEntity, err)

		assert.NotNil(t, got)
		assert.True(t, IsUnprocessableEntityError(got))
		assert.Equal(t, http.StatusUnprocessableEntity, got.Code())
		assert.Equal(t, errMsgUnprocessableEntity, got.Message())
		assert.EqualError(t, got, "original error")
	})
}

func Test_IsUnprocessableEntityError(t *testing.T) {
	t.Run("ok - return true on NewUnprocessableEntityError", func(t *testing.T) {
		got := NewUnprocessableEntityError(errMsgUnprocessableEntity)
		assert.True(t, IsUnprocessableEntityError(got))
	})

	t.Run("ok - return true on custom error with code 422", func(t *testing.T) {
		got := NewCustomError(http.StatusUnprocessableEntity, errMsgUnprocessableEntity)
		assert.True(t, IsUnprocessableEntityError(got))
	})
}
