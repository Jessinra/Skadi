package domain

import (
	"strings"

	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/validation"
	"gorm.io/gorm"
)

type User struct {
	BaseModel

	ID                uint64 `gorm:"primary_key"`
	Name              string
	Email             string `gorm:"uniqueIndex"`
	PasswordHashed    string `gorm:"password"`
	PhoneNumber       string
	ProfilePictureURL string

	UserCurrencyConfig `gorm:"embedded"`
}

type UserCurrencyConfig struct {
	CurrencyMain string
	CurrencySub  *string
}

func (User) TableName() string {
	return "users"
}

func (u *User) Validate() error {
	if err := validation.ValidateStruct(u,
		validation.Field(&u.Name, validation.Required),
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.PasswordHashed, validation.Required),
	); err != nil {
		errMsg := strings.ReplaceAll(err.Error(), ".", "")
		return errors.NewUnprocessableEntityError(errMsg)
	}

	return nil
}

// BeforeCreate gorm callback before insert into DB.
func (u *User) BeforeCreate(_ *gorm.DB) error {
	return u.Validate()
}
