package repositories

import (
	"strings"

	"gitlab.com/trivery-id/skadi/utils/errors"
	"gorm.io/gorm"
)

// NewRepositoryError handle and wrap database error and gorm error into common app error, allowing custom message to be passed.
func NewRepositoryError(msg string, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.NewNotFoundError(msg, err)
	}

	// postgres foreign_key_violation
	if strings.Contains(err.Error(), "SQLSTATE 23503") {
		return errors.NewUnprocessableEntityError(msg, err)
	}

	// postgres unique_violation
	if strings.Contains(err.Error(), "SQLSTATE 23505") {
		return errors.NewResourceAlreadyExistsError(msg, err)
	}

	// validation error
	if errors.IsUnprocessableEntityError(err) {
		return err
	}

	return errors.NewPostgresDatabaseError(msg, err)
}
