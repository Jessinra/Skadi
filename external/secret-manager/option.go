package secret

type Option func(o *Options)

type Options struct {
	Location string
	Version  string
}

func ParseOptions(opts ...Option) *Options {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithLocation(location string) Option {
	return func(o *Options) {
		o.Location = location
	}
}

func WithVersion(version string) Option {
	return func(o *Options) {
		o.Version = version
	}
}
