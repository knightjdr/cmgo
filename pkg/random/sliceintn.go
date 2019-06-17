// Package random contains functions for operations with random effects.
package random

import (
	"math/rand"
	"time"
)

// SliceIntN will return a slice of n randomly selected values
// from the slice s.
func SliceIntN(inputSlice []int, n int) []int {
	rand.Seed(time.Now().UnixNano())

	randomItems := make([]int, n)
	slice := make([]int, len(inputSlice))
	copy(slice, inputSlice)
	for i := 0; i < n; i++ {
		index := rand.Intn(len(slice))
		randomItems[i] = slice[index]
		slice = append(slice[:index], slice[index+1:]...)
	}
	return randomItems
}
