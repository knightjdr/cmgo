package slice

import (
	"sort"
	"strings"
)

// SortStringsCaseInsensitive sorts a slice of strings ignoring case.
func SortStringsCaseInsensitive(inputSlice []string) []string {
	slice := make([]string, len(inputSlice))
	copy(slice, inputSlice)

	sort.Slice(slice, func(i, j int) bool { return strings.ToLower(slice[i]) < strings.ToLower(slice[j]) })

	return slice
}
