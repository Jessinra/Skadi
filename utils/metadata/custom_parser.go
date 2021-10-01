package metadata

import "context"

// LoggerContextparser is a custom metadata parser for logging.
type LoggerContextparser struct{}

// Parse get custom metadata from parnassus context and mask it properly.
func (LoggerContextparser) Parse(ctx context.Context) map[string]interface{} {
	meta := map[string]interface{}{}
	if uuid := GetUUIDFromContext(ctx); uuid != "" {
		meta["UUID"] = uuid
	}
	if userMetadata := GetUserFromContext(ctx); userMetadata != nil {
		meta["User"] = userMetadata
	}

	return meta
}
