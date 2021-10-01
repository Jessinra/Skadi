package secret

//nolint // exported: type name will be used as secret.SecretManager by other packages, and that stutters
type SecretManager interface {
	LoadSecret(secretID string, secretContainer interface{}, opts ...Option) error
}
