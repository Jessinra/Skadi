package errors_test

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/errors"
)

const errMsgDatabase = "it's a database error"

func Test_NewDatabaseError(t *testing.T) {
	t.Run("ok - created a database error", func(t *testing.T) {
		got := NewPostgresDatabaseError(errMsgDatabase)

		assert.NotNil(t, got)
		assert.True(t, IsDatabaseError(got))
		assert.Equal(t, http.StatusInternalServerError, got.Code())
		assert.EqualError(t, got, errMsgDatabase)
	})

	t.Run("ok - created a database error with original error", func(t *testing.T) {
		err := errors.New("original error")
		got := NewPostgresDatabaseError(errMsgDatabase, err)

		assert.NotNil(t, got)
		assert.True(t, IsDatabaseError(got))
		assert.Equal(t, http.StatusInternalServerError, got.Code())
		assert.Equal(t, errMsgDatabase, got.Message())
		assert.EqualError(t, got, "original error")
	})
}

func Test_IsDatabaseError(t *testing.T) {
	t.Run("ok - return true on NewPostgresDatabaseError", func(t *testing.T) {
		got := NewPostgresDatabaseError(errMsgDatabase)
		assert.True(t, IsDatabaseError(got))
	})

	t.Run("ok - return false on custom error", func(t *testing.T) {
		got := NewCustomError(http.StatusInternalServerError, errMsgDatabase)
		assert.False(t, IsDatabaseError(got))
	})
}
