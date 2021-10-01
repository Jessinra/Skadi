package errors

import (
	"net/http"

	"github.com/pkg/errors"
)

type DBType int

const (
	dbPostgres DBType = iota + 1
	dbDynamoDB
)

// NewDatabaseError create a new postgres database error instance.
func NewPostgresDatabaseError(message string, errs ...error) DatabaseError {
	return DatabaseError{
		dbType:  dbPostgres,
		message: message,
		errs:    errs,
	}
}

func NewDynamoDatabaseError(message string, errs ...error) DatabaseError {
	return DatabaseError{
		dbType:  dbDynamoDB,
		message: message,
		errs:    errs,
	}
}

func IsDatabaseError(err error) bool {
	return errors.As(err, &DatabaseError{})
}

// DatabaseError is an error that occurred because possible data structure mistakes.
type DatabaseError struct {
	dbType  DBType
	message string
	errs    []error
}

func (DatabaseError) Code() int {
	return http.StatusInternalServerError
}

func (e DatabaseError) Message() string {
	return e.message
}

func (e DatabaseError) Error() string {
	if len(e.errs) > 0 {
		return e.errs[0].Error()
	}

	return e.Message()
}
