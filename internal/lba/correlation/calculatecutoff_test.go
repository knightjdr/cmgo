package correlation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate ideal correlation cutoff", func() {
	It("should return cutoff that would yield average number of edges closest to two", func() {
		matrix := [][]float64{
			{1, -0.2, 0.71, 0.82},
			{0.65, 1, 0.95, 0.3},
			{0.71, -0.34, 1, 0.71},
			{0.8, 0.92, 0.75, 1},
		}
		Expect(calculateCutoff(matrix, 2)).To(Equal(0.71))
	})

	It("should return cutoff that would yield average number of edges closest to three", func() {
		matrix := [][]float64{
			{1, -0.2, 0.71, 0.82},
			{0.65, 1, 0.95, 0.3},
			{0.71, -0.34, 1, 0.71},
			{0.8, 0.92, 0.75, 1},
		}
		Expect(calculateCutoff(matrix, 3)).To(Equal(0.30))
	})
})
