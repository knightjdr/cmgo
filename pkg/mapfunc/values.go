// Package mapfunc contains methods for manipulating maps.
package mapfunc

// ValuesStringString returns the values from a map[string]string.
func ValuesStringString(m map[string]string) []string {
	values := make([]string, len(m))
	i := 0
	for _, value := range m {
		values[i] = value
		i++
	}

	return values
}
