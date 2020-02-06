package slice

import (
	"fmt"
	"strings"
)

// JoinInts joins a slice of ints to create a string.
func JoinInts(intSlice []int, sep string) string {
	stringSlice := make([]string, len(intSlice))

	for i, value := range intSlice {
		stringSlice[i] = fmt.Sprintf("%d", value)
	}

	return strings.Join(stringSlice, sep)
}
