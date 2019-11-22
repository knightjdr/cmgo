package nmf_test

import (
	"github.com/knightjdr/cmgo/internal/pkg/nmf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FilterBasisForRankDefiningPreys", func() {
	It("should return row numbers for the top 'genes' in each rank", func() {
		matrix := [][]float64{
			{0.4, 0.31, 0.2},
			{0.5, 0.31, 0.2},
			{0.41, 0.31, 0.5},
			{0.7, 0.75, 0.2},
			{0.1, 0.15, 0.2},
		}

		expected := [][]int{
			{1, 2, 3},
			{0, 3},
			{2},
		}

		actual := nmf.FilterBasisForRankDefiningPreys(matrix, 3, 0.25, 0.75)
		Expect(actual).To(Equal(expected), "should return matrix of top rows per rank")
	})
})
