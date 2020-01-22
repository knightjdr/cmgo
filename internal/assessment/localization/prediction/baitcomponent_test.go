package prediction

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate bait score component", func() {
	It("should return score and bait list for each prey", func() {
		baitInteractors := map[string][]string{
			"bait1": []string{"prey1", "prey2"},
			"bait2": []string{"prey3"},
			"bait3": []string{"prey1", "prey2"},
			"bait4": []string{"prey4"},
		}

		expected := map[string][]string{
			"prey1": []string{"bait1", "bait3"},
			"prey2": []string{"bait1", "bait3"},
			"prey3": []string{"bait2"},
			"prey4": []string{"bait4"},
		}

		Expect(getBaitsPerPrey(baitInteractors)).To(Equal(expected))
	})
})
