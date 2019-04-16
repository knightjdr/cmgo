package stats

// MeanFloat calculates the mean of a []float64
func MeanFloat(slice []float64) float64 {
	mean := float64(0)

	for _, value := range slice {
		mean += value
	}

	return mean / float64(len(slice))
}
