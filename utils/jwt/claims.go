package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	Claims         = jwt.Claims
	StandardClaims = jwt.StandardClaims
)

func NewStandardClaims(opts ...Option) StandardClaims {
	options := parseOptions(opts...)
	if options.ExpiresAt.IsZero() {
		options.ExpiresAt = time.Now().Add(DefaultExpireTime)
	}

	return jwt.StandardClaims{
		ExpiresAt: options.ExpiresAt.UTC().Unix(),
		IssuedAt:  time.Now().UTC().Unix(),
		Issuer:    DefaultIssuer,
	}
}
