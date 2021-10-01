package uuid

import (
	"fmt"

	"github.com/google/uuid"
	"gitlab.com/trivery-id/skadi/utils/errors"
)

func NewUUID() (_ string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.NewInternalServerError(fmt.Sprintf("%v", r))
		}
	}()

	return uuid.New().String(), nil
}
