package transmembrane

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

var _ = Describe("Find baits per prey", func() {
	It("should return map of baits that detected a prey", func() {
		saint := &saint.SAINT{
			saint.Row{Bait: "baitA", PreyGene: "preyA"},
			saint.Row{Bait: "baitA", PreyGene: "preyB"},
			saint.Row{Bait: "baitA", PreyGene: "preyD"},
			saint.Row{Bait: "baitB", PreyGene: "preyA"},
			saint.Row{Bait: "baitB", PreyGene: "preyC"},
			saint.Row{Bait: "baitC", PreyGene: "preyC"},
			saint.Row{Bait: "baitD", PreyGene: "preyC"},
			saint.Row{Bait: "baitD", PreyGene: "preyD"},
			saint.Row{Bait: "baitD", PreyGene: "preyE"},
		}
		preys := []string{"preyA", "preyB", "preyC"}

		expected := map[string]map[string]bool{
			"preyA": map[string]bool{
				"baitA": true,
				"baitB": true,
			},
			"preyB": map[string]bool{
				"baitA": true,
			},
			"preyC": map[string]bool{
				"baitB": true,
				"baitC": true,
				"baitD": true,
			},
		}
		Expect(findBaitsPerPrey(preys, saint)).To(Equal(expected))
	})
})
