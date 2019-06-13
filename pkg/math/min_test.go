package math_test

import (
	customMath "github.com/knightjdr/cmgo/pkg/math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Min", func() {
	It("should return the minimum of two integers", func() {
		comparisons := [][]int{
			{1, 2},
			{4, 2},
			{8, 8},
		}
		expected := []int{1, 2, 8}
		for i, comparison := range comparisons {
			Expect(customMath.MinInt(comparison[0], comparison[1])).To(Equal(expected[i]))
		}
	})
})
