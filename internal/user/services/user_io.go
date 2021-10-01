package services

import (
	"strings"

	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/validation"
)

type RegisterUserInput struct {
	Name     string
	Email    string
	Password string
}

func (in *RegisterUserInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.Name, validation.Required),
		validation.Field(&in.Email, validation.Required, validation.IsEmailFormat),
		validation.Field(&in.Password, validation.Required, validation.IsStrongPassword),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}

type UpdateUserInput struct {
	ID                uint64
	Name              *string
	PhoneNumber       *string
	ProfilePictureURL *string

	CurrencyMain *string
	CurrencySub  *string
}

func (in *UpdateUserInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.ID, validation.Required),
		validation.Field(&in.PhoneNumber, validation.IsE164),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}

type UpdateUserPasswordInput struct {
	ID          uint64
	Password    string
	NewPassword string
}

func (in *UpdateUserPasswordInput) Validate() error {
	if err := validation.ValidateStruct(in,
		validation.Field(&in.ID, validation.Required),
		validation.Field(&in.Password, validation.Required),
		validation.Field(&in.NewPassword, validation.Required, validation.IsStrongPassword),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewBadRequestError(errMsg)
	}

	return nil
}
