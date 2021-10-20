package services

var fileService *FileService

func InitServices() error {
	userSvc, err := NewFileService()
	if err != nil {
		return err
	}

	fileService = userSvc
	return nil
}

func InitServiceDependencies() error {
	return fileService.InitDependencies()
}

func ValidateServices() error {
	return fileService.Validate()
}

func GetFileService() *FileService {
	return fileService
}
