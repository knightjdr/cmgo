package prediction

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate bait score component", func() {
	It("should return score and bait list for each prey", func() {
		inputFiles := fileContent{
			baitInteractors: map[string][]string{
				"bait1": []string{"prey1", "prey2"},
				"bait2": []string{"prey3"},
				"bait3": []string{"prey1", "prey2"},
				"bait4": []string{"prey4"},
			},
			predictions: map[string]int{
				"prey1": 2,
				"prey2": 1,
				"prey3": 2,
				"prey4": 1,
			},
		}
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
		baitsPerPrey := map[string][]string{
			"prey1": []string{"bait1", "bait3"},
			"prey2": []string{"bait1", "bait3"},
			"prey3": []string{"bait2"},
			"prey4": []string{"bait4"},
		}

		expected := &preyBaitScore{
			"prey1": &baitScoreComponents{
				interactorBaits:  []string{"bait1", "bait3"},
				organelleBaits:   []string{"bait3"},
				scoreOrganelle:   1.00000,
				scoreSpecificity: 0.50000,
			},
			"prey2": &baitScoreComponents{
				interactorBaits:  []string{"bait1", "bait3"},
				organelleBaits:   []string{"bait1", "bait3"},
				scoreOrganelle:   0.50000,
				scoreSpecificity: 1.00000,
			},
			"prey3": &baitScoreComponents{
				interactorBaits:  []string{"bait2"},
				organelleBaits:   []string{},
				scoreOrganelle:   0.00000,
				scoreSpecificity: 0.00000,
			},
			"prey4": &baitScoreComponents{
				interactorBaits:  []string{"bait4"},
				organelleBaits:   []string{"bait4"},
				scoreOrganelle:   0.25000,
				scoreSpecificity: 1.00000,
			},
		}
		Expect(calculateBaitScoreComponent(inputFiles, baitCompartments, baitsPerPrey)).To(Equal(expected))
	})
})
