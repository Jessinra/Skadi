package metadata

import (
	"context"

	"gitlab.com/trivery-id/skadi/utils/uuid"
)

type ctxKey string

var (
	ctxUUIDKey = ctxKey("uuid")
	ctxUserKey = ctxKey("user")
)

type User struct {
	ID                uint64
	Name              string
	Email             string
	PhoneNumber       string
	ProfilePictureURL string
	CurrencyMain      string
	CurrencySub       *string
}

func NewContextWithUUID(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxUUIDKey, uuid.NewUUID())
}

func GetUUIDFromContext(ctx context.Context) string {
	id, _ := ctx.Value(ctxUUIDKey).(string)
	return id
}

func NewContextFromUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, ctxUserKey, user)
}

func IsAuthenticated(ctx context.Context) bool {
	user := GetUserFromContext(ctx)
	return user != nil && user.ID != 0
}

func GetUserFromContext(ctx context.Context) *User {
	if user, ok := ctx.Value(ctxUserKey).(User); ok {
		return &user
	}

	return nil
}
