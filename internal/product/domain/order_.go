package domain

import (
	"strings"
	"time"

	"gitlab.com/trivery-id/skadi/internal/product/enums"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/ptr"
	"gitlab.com/trivery-id/skadi/utils/validation"
)

type Order struct {
	BaseModel

	ID          uint64
	RequesterID uint64
	ShopperID   uint64
	ProductID   uint64

	Quantity int
	Unit     string
	Notes    string

	PriceID uint64
	Price   ProductPrice `gorm:"->;foreignKey:PriceID"`

	Deal  OrderDeal  `gorm:"embedded"`
	State OrderState `gorm:"embedded"`

	Cancellations []OrderCancellation `gorm:"foreignKey:OrderID"`
}

func (Order) TableName() string {
	return "orders"
}

func (o *Order) AcceptedBy(userID uint64) error {
	if o.RequesterID == userID {
		return errors.NewBadRequestError("its your own order")
	}

	o.ShopperID = userID
	return o.UpdateLastState(enums.StateAccepted)
}

func (o *Order) DroppedBy(userID uint64, reason string) error {
	if o.ShopperID != userID {
		return errors.NewBadRequestError("you are not the shopper")
	}

	o.ShopperID = 0
	o.State = NewOrderState()
	o.Cancellations = append(o.Cancellations, OrderCancellation{
		OrderID:   o.ID,
		ShopperID: o.ShopperID,
		Reason:    reason,
	})

	return nil
}

func (o *Order) UpdateLastState(state enums.OrderState) error {
	if err := ValidateStateTransitions(o.State.LastState, state); err != nil {
		return err
	}

	o.State.LastState = state

	switch state {
	case enums.StateAccepted:
		o.State.TimeOrderAccepted = ptr.Time(time.Now())
	case enums.StatePurchased:
		o.State.TimeOrderPurchased = ptr.Time(time.Now())
	case enums.StateOnTheWay:
		o.State.TimeOrderOnTheWay = ptr.Time(time.Now())
	case enums.StateDelivered:
		o.State.TimeOrderDelivered = ptr.Time(time.Now())
	case enums.StateReviewed:
		o.State.TimeOrderReviewed = ptr.Time(time.Now())
	case enums.StateCompleted:
		o.State.TimeOrderCompleted = ptr.Time(time.Now())
	}

	return nil
}

func (o *Order) IsDeletable() bool {
	return o.State.LastState == enums.StateCreated ||
		o.State.LastState == enums.StateAccepted
}

type OrderDeal struct {
	Location   string // TODO: use address
	Time       time.Time
	Method     string
	IncludeBox bool
}

func (d *OrderDeal) Validate() error {
	if err := validation.ValidateStruct(d,
		validation.Field(&d.Location, validation.Required),
		validation.Field(&d.Time, validation.Required),
		validation.Field(&d.Method, validation.Required),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewUnprocessableEntityError(errMsg)
	}

	return nil
}

type OrderCancellation struct {
	CreatedAt time.Time
	ID        uint64
	OrderID   uint64
	ShopperID uint64
	Reason    string
}

func (OrderCancellation) TableName() string {
	return "order_cancellations"
}
