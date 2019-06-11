package stats

import "math"

// MinFloat calculates the minimum value in a []float64
func MinFloat(slice []float64) float64 {
	min := math.MaxFloat64
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	return min
}
