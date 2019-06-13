package math

import "math"

// MaxSliceFloat calculates the maximum value in a []float64
func MaxSliceFloat(slice []float64) float64 {
	max := -math.MaxFloat64
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}
