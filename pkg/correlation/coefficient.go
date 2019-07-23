// Package correlation calculates the correlation coefficient between two slices.
package correlation

// Coefficient calculates the correlation coefficient between two slices.
// The slices will only be compared up to the length of the shorter slice.
// Currently, only the Pearson method is implemented.
func Coefficient(x, y []float64, method string) float64 {
	n := len(x)
	if len(y) < n {
		n = len(y)
	}

	switch method {
	case "Pearson":
		return Pearson(x[:n], y[:n])
	}
	return 0
}
