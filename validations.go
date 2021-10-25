package circleci

func validString(v *string) bool {
	return v != nil && *v != ""
}
