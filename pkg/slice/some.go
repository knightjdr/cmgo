package slice

// SomeFloat evaluates a slice of floats []float64 and returns true
// if any member passes the provided function.
func SomeFloat(slice []float64, f func(float64) bool) bool {
	for _, value := range slice {
		if f(value) {
			return true
		}
	}
	return false
}
