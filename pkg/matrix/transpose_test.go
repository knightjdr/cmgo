package matrix_test

import (
	"github.com/knightjdr/cmgo/pkg/matrix"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transpose", func() {
	It("should tranpose matrix", func() {
		x := [][]float64{
			{5, 2, 14.3, 2.1},
			{23, 17.8, 0, 0.4},
			{10, 0, 7, 15.9},
		}
		expected := [][]float64{
			{5, 23, 10},
			{2, 17.8, 0},
			{14.3, 0, 7},
			{2.1, 0.4, 15.9},
		}
		Expect(matrix.Transpose(x)).To(Equal(expected))
	})
})
