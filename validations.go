package circleci

func validString(v *string) bool {
	return v != nil && *v != ""
}

func validCheckoutKeyType(v *checkoutKeyType) bool {
	return v != nil && *v != checkoutKeyType("")
}
