package stats

// Occurrence counts the number of times a string occurs in a slice.
func Occurrence(slice []string) map[string]int {
	occurrence := make(map[string]int, 0)
	for _, value := range slice {
		occurrence[value]++
	}
	return occurrence
}
