package s3

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/awserr"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
	"gitlab.com/trivery-id/skadi/utils/errors"
)

func NewAWSS3Error(err error) error {
	var awsErr awserr.Error
	if ok := errors.As(err, &awsErr); !ok {
		errMessage := fmt.Sprintf("AWS S3: %s", err.Error())
		return errors.NewCustomError(http.StatusInternalServerError, errMessage, err)
	}

	errCode := http.StatusInternalServerError

	// Some used errors mapping
	// Check out other errors here: https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#pkg-constants
	switch awsErr.Code() {
	case awsS3.ErrCodeBucketAlreadyExists:
		errCode = http.StatusConflict
	case awsS3.ErrCodeBucketAlreadyOwnedByYou:
		errCode = http.StatusUnprocessableEntity
	case awsS3.ErrCodeInvalidObjectState:
		errCode = http.StatusUnprocessableEntity
	case awsS3.ErrCodeNoSuchBucket:
		errCode = http.StatusNotFound
	case awsS3.ErrCodeNoSuchKey:
		errCode = http.StatusNotFound
	case awsS3.ErrCodeNoSuchUpload:
		errCode = http.StatusNotFound
	case awsS3.ErrCodeObjectAlreadyInActiveTierError:
		errCode = http.StatusUnprocessableEntity
	case awsS3.ErrCodeObjectNotInActiveTierError:
		errCode = http.StatusUnprocessableEntity
	case "NotFound":
		errCode = http.StatusNotFound
	}

	errMessage := awsErr.Message()
	if errMessage == "" {
		errMessage = awsErr.Error()
	}

	return errors.NewCustomError(errCode, fmt.Sprintf("AWS S3: %s", errMessage), awsErr)
}
