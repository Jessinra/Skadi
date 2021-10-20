package enums_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/internal/product/enums"
)

func TestValidatedOrderState(t *testing.T) {
	t.Run("ok - valid state", func(t *testing.T) {
		err := ValidatedOrderState("CREATED")
		assert.Nil(t, err)
	})

	t.Run("err - invalid state", func(t *testing.T) {
		err := ValidatedOrderState("PENDING")
		assert.EqualError(t, err, "invalid order state 'PENDING': must be one of [CREATED, ACCEPTED, PURCHASED, ON_THE_WAY, DELIVERED, REVIEWED, COMPLETED]")
	})
}
