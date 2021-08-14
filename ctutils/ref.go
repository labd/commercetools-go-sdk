package ctutils

func StringRef(value string) *string {
	return &value
}

func BoolRef(value bool) *bool {
	return &value
}

func IntRef(value int) *int {
	return &value
}
