package slice

// Diff returns the different elements between two string slices
func Diff(sliceA, sliceB []string) []string {
	allElements := append(sliceA, sliceB...)
	dictA := Dict(sliceA)
	dictB := Dict(sliceB)

	diff := make([]string, 0)
	for _, element := range allElements {
		if _, ok := dictA[element]; !ok {
			diff = append(diff, element)
		} else if _, ok := dictB[element]; !ok {
			diff = append(diff, element)
		}
	}

	return diff
}
