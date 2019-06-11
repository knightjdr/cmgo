package slice

import "strconv"

// ConvertStringToFloat converts a slice of strings to a slice of float64.
func ConvertStringToFloat(slice []string) []float64 {
	converted := make([]float64, len(slice))
	for i, str := range slice {
		converted[i], _ = strconv.ParseFloat(str, 64)
	}
	return converted
}

// ConvertStringToInt converts a slice of strings to a slice of ints.
func ConvertStringToInt(slice []string) []int {
	converted := make([]int, len(slice))
	for i, str := range slice {
		converted[i], _ = strconv.Atoi(str)
	}
	return converted
}
