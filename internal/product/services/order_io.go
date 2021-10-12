package services

import (
	"strings"
	"time"

	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/validation"
)

type CreateNewOrderInput struct {
	ProductID uint64
	PriceID   uint64

	Quantity int
	Unit     string
	Notes    string

	Deal CreateNewOrderDealInput
}

type CreateNewOrderDealInput struct {
	Location   interface{}
	Time       time.Time
	Method     string
	IncludeBox bool
}

func (in *CreateNewOrderInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.ProductID, validation.Required),
		validation.Field(&in.PriceID, validation.Required),
		validation.Field(&in.Quantity, validation.Required),
		validation.Field(&in.Unit, validation.Required),
		validation.Field(&in.Deal, validation.Required),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}

type TakeOrderInput struct {
	OrderID uint64
}

func (in *TakeOrderInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.OrderID, validation.Required),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}

type DropOrderInput struct {
	OrderID uint64
	Reason  string
}

func (in *DropOrderInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.OrderID, validation.Required),
		validation.Field(&in.Reason, validation.Length(50, 400)), // nolint
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}

type GetOrderInput struct {
	OrderID uint64
}

func (in *GetOrderInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.OrderID, validation.Required),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}

type GetAllOrdersInput struct {
	Limit  *int
	Offset *int
}

func (in *GetAllOrdersInput) Validate() error {
	return nil
}

type DeleteOrderInput struct {
	OrderID uint64
}

func (in *DeleteOrderInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.OrderID, validation.Required),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}
