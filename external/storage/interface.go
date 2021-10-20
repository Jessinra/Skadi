//go:generate mockgen --build_flags=--mod=mod -package=mocks -destination=mocks/IUploader.go . IClient

package storage

import (
	"context"
	"io"
)

type IUploader interface {
	Upload(ctx context.Context, in UploadInput) (url string, err error)
}

type UploadInput struct {
	Key             string
	File            io.Reader
	AllowPublicRead bool
}
