package slice

// UniqueStrings returns only the unique values in a slice of strings
func UniqueStrings(slice []string) []string {
	dict := Dict(slice)

	unique := make([]string, len(dict))
	i := 0
	for key := range dict {
		unique[i] = key
		i++
	}
	return unique
}
