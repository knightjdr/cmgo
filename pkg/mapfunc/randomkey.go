package mapfunc

// RandomKey gets a random key from a map[string]string
func RandomKey(m map[string]string) string {
	for key := range m {
		return key
	}
	return ""
}
