// Package sort contains sorting functions
package sort

import "sort"

// KeyValueIntFloat pairs
type KeyValueIntFloat struct {
	Key   int
	Value float64
}

// KeyValueStringInt pairs
type KeyValueStringInt struct {
	Key   string
	Value int
}

// ByMapValueIntFloat64 sorts a map by float values and returns the sort order.
func ByMapValueIntFloat64(inputMap map[int]float64, direction string) []KeyValueIntFloat {
	var sortPairs []KeyValueIntFloat
	for key, value := range inputMap {
		sortPairs = append(sortPairs, KeyValueIntFloat{key, value})
	}

	if direction == "ascending" {
		sort.Slice(sortPairs, func(i, j int) bool {
			return sortPairs[i].Value < sortPairs[j].Value
		})
	} else {
		sort.Slice(sortPairs, func(i, j int) bool {
			return sortPairs[i].Value > sortPairs[j].Value
		})
	}

	return sortPairs
}

// ByMapValueStringInt sorts a map by integer values and returns the sort order.
func ByMapValueStringInt(inputMap map[string]int, direction string) []KeyValueStringInt {
	var sortPairs []KeyValueStringInt
	for key, value := range inputMap {
		sortPairs = append(sortPairs, KeyValueStringInt{key, value})
	}

	if direction == "ascending" {
		sort.Slice(sortPairs, func(i, j int) bool {
			return sortPairs[i].Value < sortPairs[j].Value
		})
	} else {
		sort.Slice(sortPairs, func(i, j int) bool {
			return sortPairs[i].Value > sortPairs[j].Value
		})
	}

	return sortPairs
}
