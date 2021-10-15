package domain

import (
	"gitlab.com/trivery-id/skadi/utils/jwt"
	"gitlab.com/trivery-id/skadi/utils/metadata"
)

type UserClaims struct {
	User metadata.User `json:"user"`
	jwt.StandardClaims
}

type UserRefreshTokenClaims struct {
	UserID uint64 `json:"user_id"`

	jwt.StandardClaims
}
