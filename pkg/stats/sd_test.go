package stats_test

import (
	"github.com/knightjdr/cmgo/pkg/stats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SD", func() {
	It("should generate the standard deviation", func() {
		slices := [][]float64{
			{1, 2, 3, 4},
			{8, 2, 17, 34},
			{5.5, 7.2, 100.9, 3.35, 53.9},
		}
		expected := []float64{
			1.290994449,
			13.93735986,
			42.84095004,
		}
		for i, slice := range slices {
			Expect(stats.SDFloat(slice)).To(BeNumerically("~", expected[i], 0.00000001))
		}
	})
})
