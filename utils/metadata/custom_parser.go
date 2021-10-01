package metadata

import (
	"context"

	"gitlab.com/trivery-id/skadi/utils/masking"
)

// AlpsContextParser is a custom context parser for alps-trails.
type AlpsContextParser struct{}

// Parse get custom metadata from parnassus context.
func (AlpsContextParser) Parse(ctx context.Context) map[string]interface{} {
	meta := map[string]interface{}{}
	if userMetadata := GetUserFromContext(ctx); userMetadata != nil {
		meta["UserID"] = userMetadata.ID
		meta["ClientID"] = userMetadata.ClientID
		meta["Username"] = userMetadata.FullName
		meta["Email"] = userMetadata.Email
	}
	if apiKeyMetadata := GetAPIKeyFromContext(ctx); apiKeyMetadata != nil {
		meta["APIKey"] = apiKeyMetadata.Key
	}
	if clientMetadata := GetClientFromContext(ctx); clientMetadata != nil {
		meta["ClientID"] = clientMetadata.ID
	}

	return meta
}

// LoggerContextparser is a custom metadata parser for logging.
type LoggerContextparser struct{}

// Parse get custom metadata from parnassus context and mask it properly.
func (LoggerContextparser) Parse(ctx context.Context) map[string]interface{} {
	meta := map[string]interface{}{}
	if uuid := GetUUIDFromContext(ctx); uuid != "" {
		meta["UUID"] = uuid
	}
	if userMetadata := GetUserFromContext(ctx); userMetadata != nil {
		userMetadata.Email = masking.Email(userMetadata.Email)
		userMetadata.FirstName = masking.Name(userMetadata.FirstName)
		userMetadata.LastName = masking.Name(userMetadata.LastName)
		userMetadata.FullName = masking.Name(userMetadata.FullName)
		meta["User"] = userMetadata
	}
	if apiKeyMetadata := GetAPIKeyFromContext(ctx); apiKeyMetadata != nil {
		meta["APIKey"] = masking.Center(apiKeyMetadata.Key, masking.Half)
	}
	if clientMetadata := GetClientFromContext(ctx); clientMetadata != nil {
		meta["ClientID"] = clientMetadata.ID
	}

	return meta
}
