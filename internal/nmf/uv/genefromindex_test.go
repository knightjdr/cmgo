package uv

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Genes from index", func() {
	It("should convert row indices to row names", func() {
		nonCharacterizingRowIndicies := [][]int{
			{1, 2, 3},
			{2, 4, 5},
			{0, 5, 4},
		}
		rowNames := []string{"a", "b", "c", "d", "e", "f"}

		expected := [][]string{
			{"b", "c", "d"},
			{"c", "e", "f"},
			{"a", "f", "e"},
		}
		Expect(geneFromIndex(nonCharacterizingRowIndicies, rowNames)).To(Equal(expected))
	})
})
