package stats

import "math"

// MinFloatSlice calculates the minimum value in a []float64
func MinFloatSlice(slice []float64) float64 {
	min := math.MaxFloat64
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	return min
}
