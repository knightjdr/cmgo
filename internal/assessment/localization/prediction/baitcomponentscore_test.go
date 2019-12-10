package prediction

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate bait score component", func() {
	It("should return score and bait list for each prey", func() {
		baitCompartments := baitInformation{
			compartmentCounts: map[int]int{
				1: 4,
				2: 1,
			},
			localizations: map[string][]int{
				"bait1": []int{1},
				"bait2": []int{1},
				"bait3": []int{1, 2},
				"bait4": []int{1},
			},
		}
		baitInteractors := map[string][]string{
			"bait1": []string{"prey1", "prey2"},
			"bait2": []string{"prey3"},
			"bait3": []string{"prey1", "prey2"},
			"bait4": []string{"prey4"},
		}
		predictions := map[string]int{
			"prey1": 2,
			"prey2": 1,
			"prey3": 2,
			"prey4": 1,
		}

		expected := &preyBaitScore{
			"prey1": &baitScoreComponents{
				baits: []string{"bait3"},
				score: 1.00000,
			},
			"prey2": &baitScoreComponents{
				baits: []string{"bait1", "bait3"},
				score: 0.50000,
			},
			"prey3": &baitScoreComponents{
				baits: []string{},
				score: 0.00000,
			},
			"prey4": &baitScoreComponents{
				baits: []string{"bait4"},
				score: 0.25000,
			},
		}
		Expect(calculateBaitScoreComponent(predictions, baitCompartments, baitInteractors)).To(Equal(expected))
	})
})
