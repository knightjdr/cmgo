package correlation_test

import (
	"fmt"

	"github.com/knightjdr/cmgo/pkg/correlation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pearson", func() {
	It("should return correlation coefficiant", func() {
		x := [][]float64{
			{-0.5, 0, 0.5, 0.25},
			{-0.5, 0, 0.5, 0.25},
			{0.6, 0.3, 0.1, -.34},
		}
		y := [][]float64{
			{-1, 0, 1, 0.5},
			{0.5, 0, -0.5, -0.25},
			{1, -0.2, -0.85, 0.33},
		}
		expected := []float64{1, -1, 0.319}
		for i := range expected {
			Expect(correlation.Pearson(x[i], y[i])).To(BeNumerically("~", expected[i], 0.0001), fmt.Sprintf("should return %f for x an y index %d", expected[i], i))
		}
	})
})
