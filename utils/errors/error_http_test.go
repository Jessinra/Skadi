package errors_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/errors"
)

func TestGetHTTPStatus(t *testing.T) {
	t.Run("ok - got http status from a bad request error", func(t *testing.T) {
		err := NewBadRequestError("")

		got := GetHTTPStatus(err)
		assert.Equal(t, http.StatusBadRequest, got)
	})

	t.Run("ok - got http status from a custom error", func(t *testing.T) {
		err := NewCustomError(400, "")

		got := GetHTTPStatus(err)
		assert.Equal(t, http.StatusBadRequest, got)
	})

	t.Run("ok - got http status from a database error", func(t *testing.T) {
		err := NewPostgresDatabaseError("")

		got := GetHTTPStatus(err)
		assert.Equal(t, http.StatusInternalServerError, got)
	})

	t.Run("ok - got http status from a forbidden error", func(t *testing.T) {
		err := NewForbiddenError("")

		got := GetHTTPStatus(err)
		assert.Equal(t, http.StatusForbidden, got)
	})

	t.Run("ok - got http status from an internal server error", func(t *testing.T) {
		err := NewInternalServerError("")

		got := GetHTTPStatus(err)
		assert.Equal(t, http.StatusInternalServerError, got)
	})

	t.Run("ok - got http status from a not found error", func(t *testing.T) {
		err := NewNotFoundError("")

		got := GetHTTPStatus(err)
		assert.Equal(t, http.StatusNotFound, got)
	})

	t.Run("ok - got http status from a not implemented error", func(t *testing.T) {
		err := NewNotImplementedError("")

		got := GetHTTPStatus(err)
		assert.Equal(t, http.StatusNotImplemented, got)
	})

	t.Run("ok - got http status from a resource already exists error", func(t *testing.T) {
		err := NewResourceAlreadyExistsError("")

		got := GetHTTPStatus(err)
		assert.Equal(t, http.StatusConflict, got)
	})

	t.Run("ok - got http status from an unauthorized error", func(t *testing.T) {
		err := NewUnauthorizedError("")

		got := GetHTTPStatus(err)
		assert.Equal(t, http.StatusUnauthorized, got)
	})

	t.Run("ok - got http status from an unprocessable entity error", func(t *testing.T) {
		err := NewUnprocessableEntityError("")

		got := GetHTTPStatus(err)
		assert.Equal(t, http.StatusUnprocessableEntity, got)
	})
}
