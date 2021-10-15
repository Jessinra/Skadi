package errors

// list of common errors, in order to have consistent error message across all services.
var (
	ErrPermissionDenied   = NewForbiddenError("no permission to perform the requested action")
	ErrInvalidCredentials = NewUnauthorizedError("invalid credentials")
)
