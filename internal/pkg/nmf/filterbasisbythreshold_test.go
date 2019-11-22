package nmf_test

import (
	"github.com/knightjdr/cmgo/internal/pkg/nmf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FilterBasisByTreshold", func() {
	It("should return rows with at least one member passing rank minimum", func() {
		matrix := [][]float64{
			{0.4, 0.31, 0.2},
			{0.5, 0.31, 0.2},
			{0.41, 0.31, 0.5},
			{0.7, 0.75, 0.2},
			{0.1, 0.15, 0.2},
		}
		threshold := 0.5
		rows := []string{"a", "b", "c", "d", "e"}

		expectedMatrix := [][]float64{
			{0.5, 0.31, 0.2},
			{0.41, 0.31, 0.5},
			{0.7, 0.75, 0.2},
		}
		expectedRows := []string{"b", "c", "d"}

		actualMatrix, actualRows := nmf.FilterBasisByTreshold(matrix, rows, threshold)
		Expect(actualMatrix).To(Equal(expectedMatrix), "should return matrix filtered by threshold")
		Expect(actualRows).To(Equal(expectedRows), "should return rows filtered by threshold")
	})
})
