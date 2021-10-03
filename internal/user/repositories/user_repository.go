package repositories

import (
	"context"

	"gitlab.com/trivery-id/skadi/internal/user/domain"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

var (
	errAddUser    = "failed to add new user to database"
	errUpdateUser = "failed to update user in database"
	errFindUser   = "failed to find user in database"
)

func (r *UserRepository) Add(ctx context.Context, user *domain.User) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "UserRepository.Add"),
	)

	if err := r.DB.Create(&user).Error; err != nil {
		log.Error(errAddUser, zap.Error(err))
		return NewRepositoryError(errAddUser, err)
	}

	return nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "UserRepository.Update"),
		zap.Uint64("userID", user.ID),
	)

	if err := r.DB.Where("id = ?", user.ID).Save(&user).Error; err != nil {
		log.Error(errUpdateUser, zap.Error(err))
		return NewRepositoryError(errUpdateUser, err)
	}

	return nil
}

func (r *UserRepository) Find(ctx context.Context, userID uint64) (*domain.User, error) {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "UserRepository.Find"),
		zap.Uint64("userID", userID),
	)

	user := domain.User{}
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		err = NewRepositoryError(errFindUser, err)
		if !errors.IsNotFoundError(err) {
			log.Error(errFindUser, zap.Error(err))
		}

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	log := logger.GetLoggerWithCtx(ctx).With(
		zap.String("function", "UserRepository.FindByEmail"),
	)

	user := domain.User{}
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		err = NewRepositoryError(errFindUser, err)
		if !errors.IsNotFoundError(err) {
			log.Error(errFindUser, zap.Error(err))
		}

		return nil, err
	}

	return &user, nil
}
