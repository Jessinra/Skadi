package services

import (
	"fmt"
	"os"

	"gitlab.com/trivery-id/skadi/external/aws/session"
	"gitlab.com/trivery-id/skadi/external/storage"
	"gitlab.com/trivery-id/skadi/external/storage/s3"
	"gitlab.com/trivery-id/skadi/utils/errors"
)

type FileService struct {
	Uploader storage.IUploader
}

func NewFileService() (*FileService, error) {
	return &FileService{}, nil
}

func (svc *FileService) InitDependencies() error {
	defaultPublicBucket := fmt.Sprintf("trivery-%s-skadi-public", os.Getenv("ENV"))

	sess, err := session.GetDefaultSession()
	if err != nil {
		return errors.New("failed to initialize file service: invalid uploader")
	}

	svc.Uploader, err = s3.NewUploader(defaultPublicBucket, sess)
	if err != nil {
		return errors.New("failed to initialize file service: invalid uploader")
	}

	return nil
}

func (svc *FileService) Validate() error {
	if svc.Uploader == nil {
		return errors.NewUnprocessableEntityError("invalid file service, haven't set Uploader")
	}

	return nil
}
