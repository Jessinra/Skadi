//go:generate mockgen --build_flags=--mod=mod -destination=mocks/IUserRepository.go -package=mocks . IUserRepository

package services

import (
	"context"

	skadipsql "gitlab.com/trivery-id/skadi/datasources/postgres/skadi"
	"gitlab.com/trivery-id/skadi/internal/user/domain"
	"gitlab.com/trivery-id/skadi/internal/user/repositories"
	"gitlab.com/trivery-id/skadi/utils/errors"
)

type UserService struct {
	UserRepository IUserRepository
}

type IUserRepository interface {
	Add(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Find(ctx context.Context, userID uint64) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}

func NewUserService() (*UserService, error) {
	return &UserService{
		UserRepository: repositories.NewUserRepository(skadipsql.DB),
	}, nil
}

func (UserService) InitDependencies() error {
	return nil
}

func (svc *UserService) Validate() error {
	if svc.UserRepository == nil {
		return errors.NewUnprocessableEntityError("invalid user service, haven't set UserRepository")
	}

	return nil
}
