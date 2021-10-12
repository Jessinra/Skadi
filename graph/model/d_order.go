package model

import (
	"fmt"

	"gitlab.com/trivery-id/skadi/internal/product/domain"
)

func NewOrder(in *domain.Order) *Order {
	return &Order{
		ID:           in.ID,
		Requester:    &User{},
		Shopper:      &User{},
		Product:      &Product{},
		Quantity:     in.Quantity,
		Unit:         in.Unit,
		Notes:        &in.Notes,
		Price:        NewProductPrice(&in.Price),
		Deal:         NewOrderDeal(&in.Deal),
		State:        NewOrderState(&in.State),
		Cancellation: NewOrderCancellations(in.Cancellation),
	}
}

func NewOrders(in []domain.Order) []Order {
	out := []Order{}
	for i := range in {
		out = append(out, *NewOrder(&in[i]))
	}

	return out
}

func NewOrderDeal(in *domain.OrderDeal) *OrderDeal {
	return &OrderDeal{
		Location:   fmt.Sprintf("%+v", in.Location),
		Time:       in.Time,
		Method:     in.Method,
		IncludeBox: in.IncludeBox,
	}
}

func NewOrderState(in *domain.OrderState) *OrderState {
	out := OrderState(*in)
	return &out
}

func NewOrderCancellation(in domain.OrderCancellation) OrderCancellation {
	return OrderCancellation{
		CreatedAt: in.CreatedAt,
		Reason:    in.Reason,
	}
}

func NewOrderCancellations(in []domain.OrderCancellation) []OrderCancellation {
	out := []OrderCancellation{}
	for i := range in {
		out = append(out, NewOrderCancellation(in[i]))
	}

	return out
}
