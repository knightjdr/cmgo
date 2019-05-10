package slice

// HasIntersect determines if two string slices share any element.
func HasIntersect(sliceA, sliceB []string) bool {
	dictA := Dict(sliceA)
	dictB := Dict(sliceB)

	for key := range dictA {
		if _, ok := dictB[key]; ok {
			return true
		}
	}

	return false
}
