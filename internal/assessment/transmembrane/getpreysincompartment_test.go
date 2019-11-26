package transmembrane

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get preys in compartment", func() {
	It("should return slice of preys in requested compartment", func() {
		preysPerRank := map[int][]string{
			1: []string{"a", "c", "d"},
			2: []string{"b"},
			3: []string{"e", "f"},
			4: []string{"g", "h", "i"},
		}
		ranks := []string{"1", "3"}

		expected := []string{"a", "c", "d", "e", "f"}
		Expect(getPreysInCompartment(preysPerRank, ranks)).To(Equal(expected))
	})
})
