package slice

// Contains checks if a string is found in a list of strings.
func Contains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
