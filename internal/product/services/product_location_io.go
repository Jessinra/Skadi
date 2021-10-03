package services

import (
	"strings"

	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/ptr"
	"gitlab.com/trivery-id/skadi/utils/validation"
)

type CreateNewProductLocationInput struct {
	ProductID uint64

	Text     string
	Country  string
	Province string
	City     string
	Area     *string
	Street   *string
	Building *string
	Store    *string

	Longitude *float64
	Latitude  *float64
}

func (in *CreateNewProductLocationInput) FillDefault() {
	if in.Area == nil {
		in.Area = ptr.String("")
	}
	if in.Street == nil {
		in.Street = ptr.String("")
	}
	if in.Building == nil {
		in.Building = ptr.String("")
	}
	if in.Store == nil {
		in.Store = ptr.String("")
	}
	if in.Longitude == nil {
		in.Longitude = ptr.Float64(0)
	}
	if in.Latitude == nil {
		in.Latitude = ptr.Float64(0)
	}
}

func (in *CreateNewProductLocationInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.ProductID, validation.Required),
		validation.Field(&in.Text, validation.Required),
		validation.Field(&in.Country, validation.Required),
		validation.Field(&in.Province, validation.Required),
		validation.Field(&in.City, validation.Required),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	// nitpicky when
	if err := validation.Validate(*in.Latitude, validation.IsFloatLatitude); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}
	if err := validation.Validate(*in.Longitude, validation.IsFloatLongitude); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}

type UpdateProductLocationInput struct {
	ID       uint64
	Text     *string
	Province *string
	City     *string
	Area     *string
	Street   *string
	Building *string
	Store    *string

	Longitude *float64
	Latitude  *float64
}

func (in *UpdateProductLocationInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.ID, validation.Required),
		validation.Field(&in.Latitude, validation.When(in.Latitude != nil, validation.IsFloatLatitude)),
		validation.Field(&in.Longitude, validation.When(in.Longitude != nil, validation.IsFloatLongitude)),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}
