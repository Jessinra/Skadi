package services

var userService *UserService

func InitServices() error {
	userSvc, err := NewUserService()
	if err != nil {
		return err
	}

	userService = userSvc
	return nil
}

func InitServiceDependencies() error {
	return userService.InitDependencies()
}

func ValidateServices() error {
	return userService.Validate()
}

func GetUserService() *UserService {
	return userService
}
