package model

import "gitlab.com/trivery-id/skadi/internal/user/services"

func NewAuthToken(tokens *services.GenerateAuthTokensOutput) *AuthTokens {
	return &AuthTokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
}
