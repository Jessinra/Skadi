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

func GetUserFromContext(ctx context.Context) *User {
	if user, ok := ctx.Value(ctxUserKey).(User); ok {
		return &user
	}

	// TODO: remove this
	return &User{
		ID:   1001,
		Name: "dummy",
	}
}
