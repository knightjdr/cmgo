package localize

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Countassociations", func() {
	It("should count the other preys associated with each prey", func() {
		saintData := []saint.Row{
			{Bait: "a", Prey: "W"},
			{Bait: "a", Prey: "X"},
			{Bait: "a", Prey: "Y"},
			{Bait: "b", Prey: "W"},
			{Bait: "b", Prey: "X"},
			{Bait: "b", Prey: "Z"},
		}
		expected := map[string]map[string]int{
			"W": map[string]int{
				"X": 2,
				"Y": 1,
				"Z": 1,
			},
			"X": map[string]int{
				"W": 2,
				"Y": 1,
				"Z": 1,
			},
			"Y": map[string]int{
				"W": 1,
				"X": 1,
			},
			"Z": map[string]int{
				"W": 1,
				"X": 1,
			},
		}
		Expect(countAssociations(saintData)).To(Equal(expected))
	})
})
