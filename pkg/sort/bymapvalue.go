// Package sort contains sorting functions
package sort

import "sort"

// KeyValue pairs
type KeyValue struct {
	Key   string
	Value int
}

// ByMapValueInt sorts a map by integer values and returns the sort order.
// The direction determines to sort "ascending" or "descending"
func ByMapValueInt(inputMap map[string]int, direction string) []KeyValue {
	var sortPairs []KeyValue
	for key, value := range inputMap {
		sortPairs = append(sortPairs, KeyValue{key, value})
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
