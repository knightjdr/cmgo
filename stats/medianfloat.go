package stats

import (
	"sort"
)

// MedianFloat calculates the median of a []float64
func MedianFloat(inputSlice []float64) float64 {
	// Copy and sort slice.
	slice := make([]float64, len(inputSlice))
	copy(slice, inputSlice)
	sort.Float64s(slice)

	n := len(slice)
	if len(slice)%2 == 0 {
		return (slice[n/2] + slice[(n/2)-1]) / 2
	}
	return slice[(n-1)/2]
}
