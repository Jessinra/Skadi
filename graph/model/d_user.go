package model

import "gitlab.com/trivery-id/skadi/internal/user/domain"

func NewUser(in *domain.User) *User {
	return &User{
		ID:        in.ID,
		CreatedAt: in.CreatedAt,

		Name:              in.Name,
		Email:             in.Email,
		PhoneNumber:       in.PhoneNumber,
		ProfilePictureURL: in.ProfilePictureURL,
	}
}
