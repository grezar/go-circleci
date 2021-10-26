package circleci

// String returns a pointer to the given string.
func String(v string) *string {
	return &v
}

// Bool returns a pointer to the given bool.
func Bool(v bool) *bool {
	return &v
}
