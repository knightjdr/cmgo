package math_test

import (
	customMath "github.com/knightjdr/cmgo/pkg/math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MinSlice", func() {
	It("should return the minimum in a slice of floats", func() {
		slice := []float64{0.5, 0.3, 0.31, 4.1}
		Expect(customMath.MinSliceFloat(slice)).To(Equal(0.3))
	})
})
