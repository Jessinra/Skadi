package metadata

import (
	"context"
	"strings"

	"gitlab.com/trivery-id/skadi/utils/uuid"
)

type ctxKey string

var (
	ctxUUIDKey       = ctxKey("uuid")
	ctxUserKey       = ctxKey("user")
	ctxClientKey     = ctxKey("client")
	ctxRoleKey       = ctxKey("roles")
	ctxAPIKey        = ctxKey("api_key")
	ctxAuthTokensKey = ctxKey("auth_tokens")
)

type RoleMetadata map[string]string

type UserMetadata struct {
	ID        uint64
	ClientID  uint64
	AuthID    string
	Email     string
	FirstName string
	LastName  string
	FullName  string

	AllowUnmaskedContent bool
}

type ClientMetadata struct {
	ID uint64
}

type APIKeyMetadata struct {
	Key string
}

type AuthTokensMetadata struct {
	AccessToken string
	BasicToken  string
}

func NewContextWithUUID(ctx context.Context) context.Context {
	id, _ := uuid.NewUUID()
	return context.WithValue(ctx, ctxUUIDKey, id)
}

func GetUUIDFromContext(ctx context.Context) string {
	id, _ := ctx.Value(ctxUUIDKey).(string)
	return id
}

func NewContextFromUser(ctx context.Context, user UserMetadata) context.Context {
	return context.WithValue(ctx, ctxUserKey, user)
}

func GetUserFromContext(ctx context.Context) *UserMetadata {
	if user, ok := ctx.Value(ctxUserKey).(UserMetadata); ok {
		return &user
	}

	return nil
}

func IsValidUserFromContext(ctx context.Context) bool {
	if _, ok := ctx.Value(ctxUserKey).(UserMetadata); ok {
		return ok
	}

	return false
}

func NewContextFromClient(ctx context.Context, client ClientMetadata) context.Context {
	return context.WithValue(ctx, ctxClientKey, client)
}

func GetClientFromContext(ctx context.Context) *ClientMetadata {
	if client, ok := ctx.Value(ctxClientKey).(ClientMetadata); ok {
		return &client
	}

	return nil
}

func IsValidClientFromContext(ctx context.Context) bool {
	if _, ok := ctx.Value(ctxClientKey).(ClientMetadata); ok {
		return ok
	}

	return false
}

func NewContextFromRoles(ctx context.Context, role RoleMetadata) context.Context {
	return context.WithValue(ctx, ctxRoleKey, role)
}

func HasRole(ctx context.Context, role string) bool {
	roles, ok := ctx.Value(ctxRoleKey).(RoleMetadata)
	if !ok {
		return false
	}

	_, ok = roles[strings.ToUpper(role)]
	return ok
}

func NewContextFromAPIKey(ctx context.Context, apiKey APIKeyMetadata) context.Context {
	return context.WithValue(ctx, ctxAPIKey, apiKey)
}

func GetAPIKeyFromContext(ctx context.Context) *APIKeyMetadata {
	if apiKey, ok := ctx.Value(ctxAPIKey).(APIKeyMetadata); ok {
		return &apiKey
	}

	return nil
}

func NewContextFromAuthTokens(ctx context.Context, tokens AuthTokensMetadata) context.Context {
	return context.WithValue(ctx, ctxAuthTokensKey, tokens)
}

func GetAuthTokensFromContext(ctx context.Context) *AuthTokensMetadata {
	if authTokens, ok := ctx.Value(ctxAuthTokensKey).(AuthTokensMetadata); ok {
		return &authTokens
	}

	return nil
}

func IsUnmaskedContentAllowed(ctx context.Context) bool {
	user := GetUserFromContext(ctx)
	return user.AllowUnmaskedContent
}
