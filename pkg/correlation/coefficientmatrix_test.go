package correlation_test

import (
	"github.com/knightjdr/cmgo/pkg/correlation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Coefficient matrix", func() {
	It("should return correlation coefficiant matrix between rows", func() {
		x := [][]float64{
			{-0.5, 0, 0.5, 0.25},
			{-1, 0, 1, 0.5},
			{0.6, 0.3, 0.1, -.34},
		}
		actual := correlation.CoefficientMatrix(x, true, "Pearson")
		expected := [][]float64{
			{1, 1, -0.7447},
			{1, 1, -0.7447},
			{-0.7447, -0.7447, 1},
		}
		for i, row := range actual {
			for j, coefficient := range row {
				Expect(coefficient).To(BeNumerically("~", expected[i][j], 0.0001))
			}
		}
	})

	It("should return correlation coefficiant matrix between columns", func() {
		x := [][]float64{
			{-0.5, 0, 0.5, 0.25},
			{-1, 0, 1, 0.5},
			{0.6, 0.3, 0.1, -.34},
		}
		actual := correlation.CoefficientMatrix(x, false, "Pearson")
		expected := [][]float64{
			{1, 0.9522, -0.9618, -0.9999},
			{0.9522, 1, -0.8322, -0.9571},
			{-0.9618, -0.8322, 1., 0.9572},
			{-0.9999, -0.9571, 0.9572, 1},
		}
		for i, row := range actual {
			for j, coefficient := range row {
				Expect(coefficient).To(BeNumerically("~", expected[i][j], 0.0001))
			}
		}
	})
})
