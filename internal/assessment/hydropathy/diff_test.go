package hydropathy

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Difference between SAINT and BioPlex", func() {
	It("should return IDs unique to SAINT", func() {
		bioplexIDs := []string{"a", "b", "e"}
		saintIDs := []string{"nm_a", "nm_b", "nm_c", "nm_d"}
		refseqEntrez := map[string]string{
			"nm_a": "a",
			"nm_b": "b",
			"nm_c": "c",
			"nm_d": "d",
			"nm_e": "e",
		}
		expected := []string{"nm_c", "nm_d"}
		Expect(diff(saintIDs, bioplexIDs, refseqEntrez)).To(Equal(expected))
	})
})
