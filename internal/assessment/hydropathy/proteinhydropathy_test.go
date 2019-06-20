package hydropathy

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Protein hydropathy", func() {
	It("should return average hydropathy of protein sequences", func() {
		proteins := []string{"nm_a", "nm_b", "nm_c", "nm_d"}
		refseqMapping := map[string]string{
			"nm_a": "ACDPT",
			"nm_b": "NMSXC",
			"nm_d": "TTXX",
		}
		actual := proteinHydropathy(proteins, refseqMapping)
		expected := []float64{-0.3, 0.025, -0.7}
		for i := range actual {
			Expect(actual[i]).To(BeNumerically("~", expected[i], 0.00001))
		}
	})
})
