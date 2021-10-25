package circleci

// String returns a pointer to the given string.
func String(v string) *string {
	return &v
}
