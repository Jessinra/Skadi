package enums

import (
	"fmt"
	"strings"

	"gitlab.com/trivery-id/skadi/utils/errors"
)

type OrderState = string

const (
	StateCreated   = "CREATED"
	StateAccepted  = "ACCEPTED"
	StatePurchased = "PURCHASED"
	StateOnTheWay  = "ON_THE_WAY"
	StateDelivered = "DELIVERED"
	StateReviewed  = "REVIEWED"
	StateCompleted = "COMPLETED"
)

var validStates = []OrderState{
	StateCreated,
	StateAccepted,
	StatePurchased,
	StateOnTheWay,
	StateDelivered,
	StateReviewed,
	StateCompleted,
}

func ValidatedOrderState(value interface{}) error {
	s, _ := value.(string)
	for _, v := range validStates {
		if v == s {
			return nil
		}
	}

	errMsg := fmt.Sprintf("invalid order state '%s': must be one of [%s]", s, strings.Join(validStates, ", "))
	return errors.NewUnprocessableEntityError(errMsg)
}
