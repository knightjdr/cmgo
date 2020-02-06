package math_test

import (
	customMath "github.com/knightjdr/cmgo/pkg/math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Max slice", func() {
	Context("index", func() {
		It("should return the index of the maximum value in a slice of floats", func() {
			slice := []float64{0.5, 0.3, 4.1, 0.31}
			Expect(customMath.MaxIndexSliceFloat(slice)).To(Equal(2))
		})
	})

	Context("float", func() {
		It("should return the maximum value", func() {
			slice := []float64{0.5, 0.3, 0.31, 4.1}
			Expect(customMath.MaxSliceFloat(slice)).To(Equal(4.1))
		})

		It("should return zero for slice of length 0", func() {
			slice := []float64{}
			Expect(customMath.MaxSliceFloat(slice)).To(Equal(float64(0)))
		})
	})

	Context("int", func() {
		It("should return the maximum value", func() {
			slice := []int{5, 3, 7, 4}
			Expect(customMath.MaxSliceInt(slice)).To(Equal(7))
		})

		It("should return zero for slice of length 0", func() {
			slice := []int{}
			Expect(customMath.MaxSliceInt(slice)).To(Equal(0))
		})
	})
})
