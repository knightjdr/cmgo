package mapfunc

// ReverseStringString reverse the keys and values of a map[string]string
func ReverseStringString(m map[string]string) map[string]string {
	reversed := make(map[string]string, 0)

	for key, value := range m {
		reversed[value] = key
	}

	return reversed
}
