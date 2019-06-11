package slice

// Intersect returns the shared elements between two string slices
func Intersect(sliceA, sliceB []string) []string {
	dictA := Dict(sliceA)
	dictB := Dict(sliceB)

	intersection := make([]string, 0)
	for key := range dictA {
		if _, ok := dictB[key]; ok {
			intersection = append(intersection, key)
		}
	}

	return intersection
}
