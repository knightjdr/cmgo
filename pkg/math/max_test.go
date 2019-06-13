package math_test

import (
	customMath "github.com/knightjdr/cmgo/pkg/math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Max", func() {
	It("should return the maximum of two integers", func() {
		comparisons := [][]int{
			{1, 2},
			{4, 2},
			{8, 8},
		}
		expected := []int{2, 4, 8}
		for i, comparison := range comparisons {
			Expect(customMath.MaxInt(comparison[0], comparison[1])).To(Equal(expected[i]))
		}
	})
})
