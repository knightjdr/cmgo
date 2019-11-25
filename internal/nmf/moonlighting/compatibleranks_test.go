package moonlighting

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Define moonlighting ranks", func() {
	It("should create a 2D map of ranks that are are consistent with localization moonlighting", func() {
		dissimilarityMatrix := [][]float64{
			{0, 1, 1, 0, 1},
			{0, 0, 1, 1, 0},
			{0, 0, 0, 1, 1},
			{0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0},
		}

		expected := []map[int]bool{
			{1: true, 2: true, 4: true},
			{0: true, 2: true, 3: true},
			{0: true, 1: true, 3: true, 4: true},
			{1: true, 2: true, 4: true},
			{0: true, 2: true, 3: true},
		}
		Expect(defineCompatibleRanks(dissimilarityMatrix)).To(Equal(expected))
	})
})
