package services

import (
	"strings"

	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/validation"
)

type CreateNewProductPriceInput struct {
	ProductID uint64

	Currency         string
	Price            uint64
	IsPriceEstimated bool
}

func (in *CreateNewProductPriceInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.ProductID, validation.Required),
		validation.Field(&in.Currency, validation.Required),
		validation.Field(&in.Price, validation.Required),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}

type UpdateProductPriceInput struct {
	ID               uint64
	Price            *uint64
	IsPriceEstimated *bool
}

func (in *UpdateProductPriceInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.ID, validation.Required),
		validation.Field(&in.Price, validation.When(in.Price != nil, validation.GreaterOrEqual(uint64(0)))),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}
