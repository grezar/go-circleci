package circleci

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("not found")

	ErrRequiredEitherIDOrSlug       = errors.New("either organization ID or slug is required")
	ErrRequiredContextID            = errors.New("context ID is required")
	ErrRequiredContextVariableName  = errors.New("environment variable name is required")
	ErrRequiredContextVariableValue = errors.New("missing environment variable value")
)
