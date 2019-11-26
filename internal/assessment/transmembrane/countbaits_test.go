package transmembrane

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Count organelle baits per prey", func() {
	It("should return count of cytosolic and lumenal baits for each prey", func() {
		baitsPerPrey := map[string]map[string]bool{
			"preyA": map[string]bool{
				"baitA": true,
				"baitB": true,
				"baitE": true,
			},
			"preyB": map[string]bool{
				"baitA": true,
			},
			"preyC": map[string]bool{
				"baitB": true,
				"baitD": true,
				"baitE": true,
				"baitF": true,
			},
		}
		cytosolicBaits := []string{"baitA", "baitB", "baitC"}
		lumenalBaits := []string{"baitD", "baitE", "baitF"}

		expected := map[string]map[string]int{
			"preyA": map[string]int{"cytosolic": 2, "lumenal": 1},
			"preyB": map[string]int{"cytosolic": 1, "lumenal": 0},
			"preyC": map[string]int{"cytosolic": 1, "lumenal": 3},
		}
		Expect(countOrganelleBaitsPerPrey(baitsPerPrey, cytosolicBaits, lumenalBaits)).To(Equal(expected))
	})
})
