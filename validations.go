package circleci

func validString(v *string) bool {
	return v != nil && *v != ""
}

func validCheckoutKeyType(v *CheckoutKeyTypeType) bool {
	return v != nil && *v != *CheckoutKeyType("")
}
