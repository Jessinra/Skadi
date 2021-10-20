package services

import (
	"context"
	"io"

	"gitlab.com/trivery-id/skadi/external/storage"
)

type UploadInput struct {
	FileName string
	File     io.Reader
}

func (svc *FileService) Upload(ctx context.Context, in UploadInput) (string, error) {
	return svc.Uploader.Upload(ctx, storage.UploadInput{
		Key:             in.FileName,
		File:            in.File,
		AllowPublicRead: true,
	})
}
