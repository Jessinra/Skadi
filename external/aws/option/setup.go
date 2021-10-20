package option

import "github.com/aws/aws-sdk-go/aws/session"

const defaultRegion = "ap-southeast-1"

type SetupOption func(o *SetupOptions)

type SetupOptions struct {
	Region string
	Sess   *session.Session
}

func ParseSetupOptions(opts ...SetupOption) *SetupOptions {
	SetupOptions := &SetupOptions{
		Region: defaultRegion,
	}
	for _, opt := range opts {
		opt(SetupOptions)
	}

	return SetupOptions
}

func WithLocation(location string) SetupOption {
	return func(o *SetupOptions) {
		o.Region = location
	}
}

func WithSession(sess *session.Session) SetupOption {
	return func(o *SetupOptions) {
		o.Sess = sess
	}
}
