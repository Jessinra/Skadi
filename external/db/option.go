package db

type Option func(o *Options)

type Options struct {
	// MaxOpenConnection are number of maximum kept-alive connections for each DB's open connection.
	// Default value is unlimited, any value <= 0 means unlimited.
	MaxOpenConnection int
	SilentMode        bool
}

func ParseOptions(opts ...Option) *Options {
	options := &Options{
		MaxOpenConnection: -1, // unlimited
	}

	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithMaxOpenConnection(n int) Option {
	return func(o *Options) {
		o.MaxOpenConnection = n
	}
}

func WithSilentMode() Option {
	return func(o *Options) {
		o.SilentMode = true
	}
}
