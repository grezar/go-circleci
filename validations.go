package circleci

func validString(v *string) bool {
	return v != nil && *v != ""
}

func validCheckoutKeyType(v *CheckoutKeyTypeType) bool {
	return v != nil && *v != *CheckoutKeyType("")
}

func validArrayOfEvent(v []*Event) bool {
	for _, e := range v {
		if e == nil || *e == *EventType("") {
			return false
		}
	}
	return len(v) != 0
}

func validBool(v *bool) bool {
	return v != nil
}
