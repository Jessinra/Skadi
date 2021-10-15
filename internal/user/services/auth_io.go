package services

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token"`
}

type GenerateAuthTokensOutput struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
