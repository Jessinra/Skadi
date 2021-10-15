package writer

type Option func(o *options)

type options struct {
	EnableErrorLogging bool
}

func parseOptions(opts ...Option) *options {
	options := &options{
		EnableErrorLogging: true,
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithErrorLogging(b bool) Option {
	return func(o *options) {
		o.EnableErrorLogging = b
	}
}
