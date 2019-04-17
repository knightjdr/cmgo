// Package slice contains methods for slice conversion and modification
package slice

// Dict converts a slice into a dictionary/hash
func Dict(slice []string) map[string]bool {
	dict := make(map[string]bool, len(slice))
	for _, val := range slice {
		dict[val] = true
	}
	return dict
}
