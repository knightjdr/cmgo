// Package filter contains filtering functions for maps.
package filter

// Key filters a map and returns a slice with keys that pass the occurrence cutoff.
func Key(m map[string]int, minOccurence int) []string {
	filtered := make([]string, 0)
	for key, count := range m {
		if count >= minOccurence {
			filtered = append(filtered, key)
		}
	}
	return filtered
}
