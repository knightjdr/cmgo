package mapfunc

// RandomKeyStringString gets a random key from a map[string]string
func RandomKeyStringString(m map[string]string) string {
	for key := range m {
		return key
	}
	return ""
}
