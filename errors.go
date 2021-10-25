package circleci

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("not found")

	ErrRequiredEitherIDOrSlug = errors.New("either organization ID or slug is required")
)
