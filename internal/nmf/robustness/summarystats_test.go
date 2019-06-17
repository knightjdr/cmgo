package robustness

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Summary stats", func() {
	It("should write data to tsv file", func() {
		data := [][][]float64{
			{
				{1, 2, 3},
				{1, 2, 3},
			},
			{
				{1, 2, 3},
				{1, 2, 3},
			},
		}
		percentile := []float64{0.9, 0.8}
		actual := summaryStats(data, percentile)
		expected := [][]meanSD{
			{
				{Mean: 2, SD: 1},
				{Mean: 2, SD: 1},
			},
			{
				{Mean: 2, SD: 1},
				{Mean: 2, SD: 1},
			},
		}
		Expect(actual).To(Equal(expected))
	})
})
