package domain

import "gitlab.com/trivery-id/skadi/utils/jwt"

type UserClaims struct {
	UserID                uint64  `json:"user_id"`
	UserName              string  `json:"user_name"`
	UserEmail             string  `json:"user_email"`
	UserPhoneNumber       string  `json:"user_phone_number"`
	UserProfilePictureURL string  `json:"user_profile_picture_url"`
	CurrencyMain          string  `json:"user_currency_main"`
	CurrencySub           *string `json:"user_currency_sub"`

	jwt.StandardClaims
}

type UserRefreshTokenClaims struct {
	UserID uint64 `json:"user_id"`

	jwt.StandardClaims
}
