package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	DefaultIssuer     = "skadi"
	DefaultExpireTime = 30 * time.Minute
)

var (
	errUnexpectedSignMethod = errors.New("unexpected JWT token signing method")
	errInvalidToken         = errors.New("invalid JWT token")
)

func NewToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTSignKey())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, getKeyFunc)
	if err != nil {
		return nil, errInvalidToken
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !(ok && parsedToken.Valid) {
		return nil, errInvalidToken
	}

	return claims, nil
}

func getKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errUnexpectedSignMethod
	}

	return getJWTSignKey(), nil
}
