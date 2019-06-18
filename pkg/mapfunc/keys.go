// Package mapfunc contains methods for manipulating maps.
package mapfunc

// KeysIntFloat returns the keys of a map[int]float64.
func KeysIntFloat(m map[int]float64) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

// KeysStringBool returns the keys of a map[string]bool.
func KeysStringBool(m map[string]bool) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}
