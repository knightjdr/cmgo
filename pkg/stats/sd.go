package stats

import (
	"math"
)

// SDFloat calculates the standard deviation of []float64.
func SDFloat(slice []float64) float64 {
	if len(slice) == 1 {
		return 0
	}

	mean := MeanFloat(slice)
	sum := float64(0)
	for _, value := range slice {
		sum += math.Pow(value-mean, float64(2))
	}

	return math.Sqrt(sum / float64(len(slice)-1))
}

// SDInt calculates the standard deviation of []int.
func SDInt(slice []int) float64 {
	if len(slice) == 1 {
		return 0
	}

	mean := MeanInt(slice)
	sum := float64(0)
	for _, value := range slice {
		sum += math.Pow(float64(value)-mean, float64(2))
	}

	return math.Sqrt(sum / float64(len(slice)-1))
}
