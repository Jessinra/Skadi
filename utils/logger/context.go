package logger

import "context"

var defaultContextParser ContextParser

type ContextParser interface {
	Parse(ctx context.Context) map[string]interface{}
}

// SetDefaultContextParser set the default parser that will be used to extract metadata from the given context.
func SetDefaultContextParser(parser ContextParser) {
	defaultContextParser = parser
}
