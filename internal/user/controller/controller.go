package controller

import (
	"gitlab.com/trivery-id/skadi/internal/user/services"
	"gitlab.com/trivery-id/skadi/utils/errors"
)

var authController *AuthController

func GetAuthController() *AuthController {
	return authController
}

func InitControllers() error {
	ctrl, err := NewAuthController()
	if err != nil {
		return err
	}

	authController = ctrl
	return nil
}

func ValidateControllers() error {
	return authController.Validate()
}

type AuthController struct {
	UserService *services.UserService
}

func NewAuthController() (*AuthController, error) {
	return &AuthController{
		UserService: services.GetUserService(),
	}, nil
}

func (ctrl *AuthController) Validate() error {
	if ctrl.UserService == nil {
		return errors.NewUnprocessableEntityError("invalid auth controller, haven't set FraudScoringService")
	}

	return nil
}
