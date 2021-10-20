package s3

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gitlab.com/trivery-id/skadi/external/storage"
	"gitlab.com/trivery-id/skadi/utils/errors"
)

type Uploader struct {
	uploader   *s3manager.Uploader
	bucketName string
}

func NewUploader(bucketName string, sess *session.Session) (*Uploader, error) {
	if bucketName == "" {
		return nil, errors.NewBadRequestError("invalid bucket name")
	}

	uploader := s3manager.NewUploader(sess)
	return &Uploader{
		uploader:   uploader,
		bucketName: bucketName,
	}, nil
}

func (s *Uploader) Upload(ctx context.Context, in storage.UploadInput) (url string, err error) {
	input := &s3manager.UploadInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(in.Key),
		Body:   in.File,
	}

	if in.AllowPublicRead {
		input.ACL = aws.String("public-read")
	}

	result, err := s.uploader.UploadWithContext(ctx, input)
	if err != nil {
		return "", NewAWSS3Error(err)
	}

	return result.Location, nil
}
