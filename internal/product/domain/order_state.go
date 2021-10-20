package domain

import (
	"fmt"
	"time"

	"gitlab.com/trivery-id/skadi/internal/product/enums"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/ptr"
)

type OrderState struct {
	LastState enums.OrderState

	TimeOrderCreated   *time.Time
	TimeOrderAccepted  *time.Time
	TimeOrderPurchased *time.Time
	TimeOrderOnTheWay  *time.Time
	TimeOrderDelivered *time.Time
	TimeOrderReviewed  *time.Time
	TimeOrderCompleted *time.Time
}

func NewOrderState() OrderState {
	return OrderState{
		LastState:        enums.StateCreated,
		TimeOrderCreated: ptr.Time(time.Now()),
	}
}

var AllowedOrderStateTransitions = map[enums.OrderState]map[enums.OrderState]bool{
	enums.StateCreated: {
		enums.StateAccepted: true,
	},
	enums.StateAccepted: {
		enums.StateCreated:   true, // allow drop order by shopper
		enums.StatePurchased: true,
	},
	enums.StatePurchased: {
		enums.StateCreated:  true, // allow drop order by shopper
		enums.StateOnTheWay: true,
	},
	enums.StateOnTheWay: {
		enums.StateCreated:   true, // allow drop order by shopper
		enums.StateDelivered: true,
	},
	enums.StateDelivered: {
		enums.StateCreated:   true, // allow drop order by shopper
		enums.StateReviewed:  true,
		enums.StateCompleted: true, // skip review process
	},
	enums.StateReviewed: {
		enums.StateCreated:   true, // allow drop order by shopper
		enums.StateCompleted: true,
	},
}

func ValidateStateTransitions(from, to enums.OrderState) error {
	if from == to {
		return nil
	}

	if AllowedOrderStateTransitions[from][to] {
		return nil
	}

	errMsg := fmt.Sprintf("invalid order state transition from '%s' to '%s'", from, to)
	return errors.NewUnprocessableEntityError(errMsg)
}
