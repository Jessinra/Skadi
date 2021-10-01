package session

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	defaultSession *session.Session
	defaultRegion  = "ap-southeast-1"
)

func GetDefaultSession() (*session.Session, error) {
	if defaultSession == nil {
		sess, err := initDefaultSession()
		if err != nil {
			return nil, err
		}

		defaultSession = sess
	}

	return defaultSession, nil
}

// initDefaultSession initialize a new session and cache it for future uses.
// Ref: https://docs.aws.amazon.com/sdk-for-go/api/aws/session/
// Sessions are safe to use concurrently as long as the Session is not being modified. Sessions should be cached when possible,
// because creating a new Session will load all configuration values from the environment, and config files each time the Session is created.
func initDefaultSession() (*session.Session, error) {
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	credential := credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(defaultRegion),
		Credentials: credential,
	})
	if err != nil {
		return nil, err
	}

	return sess, nil
}
