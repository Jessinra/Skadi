package jwt

import "time"

type Option func(o *options)

type options struct {
	ExpiresAt time.Time
}

func parseOptions(opts ...Option) *options {
	options := &options{}
	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithExpiresAt(t time.Time) Option {
	return func(o *options) {
		o.ExpiresAt = t
	}
}
