package sha

import (
	"crypto/sha512"
	"encoding/base64"

	"gitlab.com/trivery-id/skadi/utils/logger"
	"go.uber.org/zap"
)

// Hash512 consistently hash input string using sha512 algorithm.
// DO NOT use this for credentials such as password; please use other more-secured algorithms like Argon2, Bcrypt, PBKDF2, etc.
func Hash512(in string) string {
	hasher := sha512.New()
	if _, err := hasher.Write([]byte(in)); err != nil {
		logger.Warn("can't write hashed string",
			zap.String("function", "sha.Hash512"),
			zap.Error(err),
		)
	}

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
