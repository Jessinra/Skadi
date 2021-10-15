package services

import (
	"context"

	"gitlab.com/trivery-id/skadi/internal/user/domain"
	"gitlab.com/trivery-id/skadi/utils/crypto/sha"
	"gitlab.com/trivery-id/skadi/utils/errors"
)

func (svc *UserService) RegisterUser(ctx context.Context, in RegisterUserInput) (*domain.User, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	_, err := svc.UserRepository.FindByEmail(ctx, in.Email)
	switch {
	case err == nil:
		return nil, errors.New("email already registered")
	case !errors.IsNotFoundError(err):
		return nil, err
	}

	user := &domain.User{
		Name:  in.Name,
		Email: in.Email,
	}
	if err := user.SetPassword(in.Password); err != nil {
		return nil, err
	}

	if err := svc.UserRepository.Add(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *UserService) GetUser(ctx context.Context, userID uint64) (*domain.User, error) {
	return svc.UserRepository.Find(ctx, userID)
}

func (svc *UserService) UpdateUser(ctx context.Context, in UpdateUserInput) (*domain.User, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	user, err := svc.UserRepository.Find(ctx, in.ID)
	if err != nil {
		return nil, err
	}

	if in.Name != nil {
		user.Name = *in.Name
	}
	if in.PhoneNumber != nil {
		user.PhoneNumber = *in.PhoneNumber
	}
	if in.ProfilePictureURL != nil {
		user.ProfilePictureURL = *in.ProfilePictureURL
	}
	if in.CurrencyMain != nil {
		user.CurrencyMain = *in.CurrencyMain
	}
	if in.CurrencySub != nil {
		user.CurrencySub = in.CurrencySub
	}
	if err := svc.UserRepository.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *UserService) UpdateUserPassword(ctx context.Context, in UpdateUserPasswordInput) error {
	if err := in.Validate(); err != nil {
		return err
	}

	user, err := svc.UserRepository.Find(ctx, in.ID)
	if err != nil {
		return err
	}

	if user.PasswordHashed != sha.Hash512(in.Password) {
		return errors.New("invalid current password")
	}

	user.PasswordHashed = sha.Hash512(in.NewPassword)
	return svc.UserRepository.Update(ctx, user)
}
