package services

import (
	"strings"

	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/ptr"
	"gitlab.com/trivery-id/skadi/utils/validation"
)

type CreateNewProductInput struct {
	Name        string
	Description string
	ImagesURLs  []string
	Weight      *string
	Dimensions  *string
	Categories  []string

	Location CreateNewProductLocationInput
	Price    CreateNewProductPriceInput
}

func (in *CreateNewProductInput) FillDefault() {
	if in.Weight == nil {
		in.Weight = ptr.String("")
	}
	if in.Dimensions == nil {
		in.Dimensions = ptr.String("")
	}
}

func (in *CreateNewProductInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.Name, validation.Required),
		validation.Field(&in.Description, validation.Required),
		validation.Field(&in.ImagesURLs, validation.Required),
		validation.Field(&in.Categories, validation.Required),
		validation.Field(&in.Location),
		validation.Field(&in.Price),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}

type GetAllProductsInput struct {
	Limit  *int
	Offset *int
}
