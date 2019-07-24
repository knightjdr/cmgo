package correlation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filter pairs", func() {
	Context("no max edge argument", func() {
		It("should return edge pairs", func() {
			corr := [][]float64{
				{1, -0.2, 0.73, 0.82},
				{-0.2, 1, 0.95, 0.3},
				{0.73, -0.95, 1, 0.71},
				{0.82, 0.3, 0.71, 1},
			}
			genes := []string{"a", "b", "c", "d"}
			expected := map[string][]edgePair{
				"a": []edgePair{
					{Target: "d", Weight: 0.82},
					{Target: "c", Weight: 0.73},
				},
				"b": []edgePair{
					{Target: "c", Weight: 0.95},
				},
				"c": []edgePair{
					{Target: "d", Weight: 0.71},
				},
				"d": []edgePair{},
			}
			Expect(filterPairs(corr, genes, 0.5, 0)).To(Equal(expected))
		})
	})

	Context("max edge argument", func() {
		It("should return edge pairs", func() {
			corr := [][]float64{
				{1, -0.2, 0.73, 0.82},
				{-0.2, 1, 0.95, 0.3},
				{0.73, -0.95, 1, 0.71},
				{0.82, 0.3, 0.71, 1},
			}
			genes := []string{"a", "b", "c", "d"}
			expected := map[string][]edgePair{
				"a": []edgePair{
					{Target: "d", Weight: 0.82},
				},
				"b": []edgePair{
					{Target: "c", Weight: 0.95},
				},
				"c": []edgePair{
					{Target: "a", Weight: 0.73},
				},
				"d": []edgePair{},
			}
			Expect(filterPairs(corr, genes, 0.5, 1)).To(Equal(expected))
		})
	})
})
