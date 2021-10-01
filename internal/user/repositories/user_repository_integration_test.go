package repositories_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gitlab.com/trivery-id/skadi/internal/user/domain"
	"gitlab.com/trivery-id/skadi/utils/errors"
	. "gitlab.com/trivery-id/skadi/utils/test"
)

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, NewUserRepositorySuite(cred))
}

func (s *UserRepositorySuite) createDummyUser() domain.User {
	s.T().Helper()

	user := domain.User{
		Name:              Random("name"),
		Email:             Random("email"),
		PasswordHashed:    Random("password"),
		PhoneNumber:       Random("phone_number"),
		ProfilePictureURL: "http://example.com/profile.png",
	}

	if err := s.TestDB.Create(&user).Error; err != nil {
		s.T().Fatal(err.Error())
	}

	return user
}

func (s *UserRepositorySuite) TestAdd() {
	s.T().Run("ok - added new user", func(t *testing.T) {
		user := domain.User{
			Name:              Random("name"),
			Email:             Random("email"),
			PasswordHashed:    Random("password"),
			PhoneNumber:       Random("phone_number"),
			ProfilePictureURL: "http://example.com/profile.png",
		}

		err := s.UserRepository.Add(context.Background(), &user)
		assert.Nil(s.T(), err)

		s.T().Run("check user created in database", func(t *testing.T) {
			got := domain.User{}
			err := s.TestDB.Where("id = ?", user.ID).Find(&got).Error
			assert.NotNil(s.T(), got)
			assert.Equal(s.T(), user.ID, got.ID)
			assert.Equal(s.T(), user.Name, got.Name)
			assert.Equal(s.T(), user.Email, got.Email)
			assert.Equal(s.T(), user.PasswordHashed, got.PasswordHashed)
			assert.Equal(s.T(), user.PhoneNumber, got.PhoneNumber)
			assert.Equal(s.T(), user.ProfilePictureURL, got.ProfilePictureURL)

			assert.Nil(s.T(), err)
		})
	})

	s.T().Run("err - duplicate email", func(t *testing.T) {
		existingUser := s.createDummyUser()
		user := domain.User{
			Name:           Random("name"),
			Email:          existingUser.Email,
			PasswordHashed: Random("password"),
		}

		err := s.UserRepository.Add(context.Background(), &user)
		assert.NotNil(s.T(), err)
		assert.True(s.T(), errors.IsResourceAlreadyExistsError(err))
		assert.Contains(s.T(), err.Error(), "email")
	})

	s.T().Run("err - empty name", func(t *testing.T) {
		user := domain.User{
			Name:           "",
			Email:          Random("email"),
			PasswordHashed: Random("password"),
		}

		err := s.UserRepository.Add(context.Background(), &user)
		assert.NotNil(s.T(), err)
		assert.True(s.T(), errors.IsUnprocessableEntityError(err))
		assert.Contains(s.T(), err.Error(), "Name")
	})

	s.T().Run("err - empty password", func(t *testing.T) {
		user := domain.User{
			Name:           Random("name"),
			Email:          Random("email"),
			PasswordHashed: "",
		}

		err := s.UserRepository.Add(context.Background(), &user)
		assert.NotNil(s.T(), err)
		assert.True(s.T(), errors.IsUnprocessableEntityError(err))
		assert.Contains(s.T(), err.Error(), "Password")
	})

	s.T().Run("err - database error", func(t *testing.T) {
		user := domain.User{
			Name:           Random("name"),
			Email:          Random("email"),
			PasswordHashed: Random("password"),
		}

		err := s.EmptyUserRepository.Add(context.Background(), &user)
		assert.NotNil(s.T(), err)
		assert.True(s.T(), errors.IsDatabaseError(err))
	})
}
