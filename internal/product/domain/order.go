package domain

import "time"

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
	Price   ProductPrice `gorm:"foreignKey:PriceID"`

	Deal  OrderDeal  `gorm:"embedded"`
	State OrderState `gorm:"embedded"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderDeal struct {
	Location   interface{}
	Date       time.Time
	Method     string
	IncludeBox bool
}

type OrderState struct {
	LastState string

	TimeOrderCreated   time.Time
	TimeOrderAccepted  time.Time
	TimeOrderOnTheWay  time.Time
	TimeOrderDelivered time.Time
	TimeOrderReviewed  time.Time
	TimeOrderCompleted time.Time
}
