package errors_test

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/errors"
)

const errMsgAlreadyExists = "it's a resource already exists error"

func Test_NewResourceAlreadyExistsError(t *testing.T) {
	t.Run("ok - created a resource already exists error", func(t *testing.T) {
		got := NewResourceAlreadyExistsError(errMsgAlreadyExists)

		assert.NotNil(t, got)
		assert.True(t, IsResourceAlreadyExistsError(got))
		assert.Equal(t, http.StatusConflict, got.Code())
		assert.EqualError(t, got, errMsgAlreadyExists)
	})

	t.Run("ok - created a resource already exists error with original error", func(t *testing.T) {
		err := errors.New("original error")
		got := NewResourceAlreadyExistsError(errMsgAlreadyExists, err)

		assert.NotNil(t, got)
		assert.True(t, IsResourceAlreadyExistsError(got))
		assert.Equal(t, http.StatusConflict, got.Code())
		assert.Equal(t, errMsgAlreadyExists, got.Message())
		assert.EqualError(t, got, "original error")
	})
}

func Test_IsResourceAlreadyExistsError(t *testing.T) {
	t.Run("ok - return true on NewResourceAlreadyExistsError", func(t *testing.T) {
		got := NewResourceAlreadyExistsError(errMsgAlreadyExists)
		assert.True(t, IsResourceAlreadyExistsError(got))
	})

	t.Run("ok - return true on custom error with code 409", func(t *testing.T) {
		got := NewCustomError(http.StatusConflict, errMsgAlreadyExists)
		assert.True(t, IsResourceAlreadyExistsError(got))
	})
}
