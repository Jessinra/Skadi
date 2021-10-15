package services

import (
	"context"
	"time"

	"gitlab.com/trivery-id/skadi/internal/user/domain"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/jwt"
)

var (
	errUnauthorized         = errors.New("invalid email/password")
	errFailedToRefreshToken = errors.New("failed to refresh token")
)

func (svc *UserService) Login(ctx context.Context, in LoginInput) (*GenerateAuthTokensOutput, error) {
	user, err := svc.UserRepository.FindByEmail(ctx, in.Email)
	if err != nil {
		return nil, errUnauthorized
	}

	if err := user.ComparePassword(in.Password); err != nil {
		return nil, errUnauthorized
	}

	return svc.GenerateAuthTokens(ctx, user.ID)
}

func (svc *UserService) RefreshToken(ctx context.Context, in RefreshTokenInput) (*GenerateAuthTokensOutput, error) {
	claims, err := jwt.ParseToken(in.RefreshToken)
	if err != nil {
		return nil, errFailedToRefreshToken
	}

	userID, _ := claims["user_id"].(uint64)
	if userID == 0 {
		return nil, errFailedToRefreshToken
	}

	return svc.GenerateAuthTokens(ctx, userID)
}

func (svc *UserService) GenerateAuthTokens(ctx context.Context, userID uint64) (*GenerateAuthTokensOutput, error) {
	var (
		defaultTokenExpiration        = 30 * time.Minute
		defaultRefreshTokenExpiration = 7 * 24 * time.Hour
	)

	user, err := svc.UserRepository.Find(ctx, userID)
	if err != nil {
		return nil, err
	}

	userClaims := domain.UserClaims{
		UserID:                user.ID,
		UserName:              user.Name,
		UserEmail:             user.Email,
		UserPhoneNumber:       user.PhoneNumber,
		UserProfilePictureURL: user.ProfilePictureURL,
		CurrencyMain:          user.CurrencyMain,
		CurrencySub:           user.CurrencySub,
		StandardClaims:        jwt.NewStandardClaims(jwt.WithExpiresAt(time.Now().Add(defaultTokenExpiration))),
	}
	jwtToken, err := jwt.NewToken(userClaims)
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := domain.UserRefreshTokenClaims{
		UserID:         user.ID,
		StandardClaims: jwt.NewStandardClaims(jwt.WithExpiresAt(time.Now().Add(defaultRefreshTokenExpiration))),
	}
	refreshToken, err := jwt.NewToken(refreshTokenClaims)
	if err != nil {
		return nil, err
	}

	return &GenerateAuthTokensOutput{
		Token:        jwtToken,
		RefreshToken: refreshToken,
	}, nil
}
