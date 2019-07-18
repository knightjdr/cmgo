package enrichment

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Associations", func() {
	It("should count the other preys associated with each prey", func() {
		saintData := []saint.Row{
			{Bait: "a", Prey: "W"},
			{Bait: "a", Prey: "X"},
			{Bait: "a", Prey: "Y"},
			{Bait: "b", Prey: "W"},
			{Bait: "b", Prey: "X"},
			{Bait: "b", Prey: "Z"},
		}
		actualBaitsPerPrey, actualPreysPerBait := associations(saintData)
		expectedBaitsPerPrey := map[string][]string{
			"W": []string{"a", "b"},
			"X": []string{"a", "b"},
			"Y": []string{"a"},
			"Z": []string{"b"},
		}
		expectedPreysPerBait := map[string][]string{
			"a": []string{"W", "X", "Y"},
			"b": []string{"W", "X", "Z"},
		}

		Expect(actualBaitsPerPrey).To(Equal(expectedBaitsPerPrey))
		Expect(actualPreysPerBait).To(Equal(expectedPreysPerBait))
	})
})
