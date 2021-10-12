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
	o.State.LastState = enums.StateAccepted
	o.State.TimeOrderAccepted = ptr.Time(time.Now())
	return nil
}

func (o *Order) Drop(userID uint64, reason string) error {
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

func NewOrderState() OrderState {
	return OrderState{
		LastState:        enums.StateCreated,
		TimeOrderCreated: ptr.Time(time.Now()),
	}
}

type OrderState struct {
	LastState string

	TimeOrderCreated   *time.Time
	TimeOrderAccepted  *time.Time
	TimeOrderPurchased *time.Time
	TimeOrderOnTheWay  *time.Time
	TimeOrderDelivered *time.Time
	TimeOrderReviewed  *time.Time
	TimeOrderCompleted *time.Time
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
