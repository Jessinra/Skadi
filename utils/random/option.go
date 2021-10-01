package random

type Option func(o *options)

type options struct {
	runes []rune
}

func parseOptions(opts ...Option) *options {
	options := &options{
		runes: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"),
	}
	for _, opt := range opts {
		opt(options)
	}

	return options
}

func WithRunes(runes []rune) Option {
	return func(o *options) {
		o.runes = runes
	}
}
