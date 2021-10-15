package jwt

import (
	"fmt"
	"os"

	"gitlab.com/trivery-id/skadi/external/secret-manager/aws"
	"gitlab.com/trivery-id/skadi/utils/logger"
)

var signKey []byte

// SetSignKey should only be used in test.
func SetSignKey(key string) {
	signKey = []byte(key)
}

func getJWTSignKey() []byte {
	if signKey == nil {
		signKey = getSignKeyFromSecretManager()
	}

	return signKey
}

func getSignKeyFromSecretManager() []byte {
	secret := struct {
		JWT struct {
			SignKey string `json:"sign_key"`
		} `json:"jwt"`
	}{}

	secretName := getAPPSecretName()
	if err := aws.NewSecretManager().LoadSecret(secretName, &secret); err != nil {
		logger.Error("Failed to get secret", err)
	}

	return []byte(secret.JWT.SignKey)
}

func getAPPSecretName() string {
	return fmt.Sprintf("skadi-%s", os.Getenv("ENV"))
}
