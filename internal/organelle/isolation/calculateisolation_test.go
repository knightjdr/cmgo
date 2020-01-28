package isolation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate isolation score", func() {
	It("should calculate score", func() {
		correlation := [][]float64{
			{1, 0.71, 0.75, 0.9},
			{0.71, 1, 0.68, 0.34},
			{0.75, 0.68, 1, 0.74},
			{0.9, 0.34, 0.74, 1},
		}
		cutoff := 0.7
		preysPerCompartment := map[int]map[int]bool{
			1: map[int]bool{0: true, 3: true},
			2: map[int]bool{1: true, 2: true},
		}

		expected := &isolationScores{
			1: &isolationScore{
				edgesOutside:       3,
				edgesWithin:        1,
				isolation:          0.25,
				nodesOutside:       []int{1, 2, 2},
				sharedCompartments: []int{0, 0, 0},
			},
			2: &isolationScore{
				edgesOutside:       3,
				edgesWithin:        0,
				isolation:          0,
				nodesOutside:       []int{0, 0, 3},
				sharedCompartments: []int{0, 0, 0},
			},
		}
		Expect(calculateIsolation(correlation, cutoff, preysPerCompartment)).To(Equal(expected))
	})
})
