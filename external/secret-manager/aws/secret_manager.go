package aws

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"gitlab.com/trivery-id/skadi/external/aws/session"
	"gitlab.com/trivery-id/skadi/external/secret-manager"
)

const (
	defaultRegion  = "ap-southeast-1"
	defaultVersion = "AWSCURRENT"
)

type awsSecretManager struct{}

func NewSecretManager() secret.SecretManager {
	return awsSecretManager{}
}

func (awsSecretManager) LoadSecret(secretID string, secretContainer interface{}, opts ...secret.Option) error {
	options := secret.ParseOptions(opts...)
	if options.Location == "" {
		options.Location = defaultRegion
	}
	if options.Version == "" {
		options.Version = defaultVersion
	}

	sess, err := session.GetDefaultSession()
	if err != nil {
		return err
	}

	secretManager := secretsmanager.New(sess,
		aws.NewConfig().WithRegion(options.Location),
	)

	result, err := secretManager.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretID),
		VersionStage: aws.String(options.Version),
	})
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(*result.SecretString), &secretContainer)
}
