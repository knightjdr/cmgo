package nmf_test

import (
	"github.com/knightjdr/cmgo/internal/pkg/nmf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Preys localizing to each rank", func() {
	It("should return preys localizing to each rank", func() {
		matrix := [][]float64{
			{0.4, 0.31, 0.2},
			{0.5, 0.31, 0.2},
			{0.41, 0.31, 0.5},
			{0.7, 0.75, 0.2},
			{0.1, 0.15, 0.2},
		}
		rows := []string{"a", "b", "c", "d", "e"}

		expected := map[int][]string{
			1: []string{"a", "b"},
			2: []string{"d"},
			3: []string{"c", "e"},
		}

		Expect(nmf.GetPreysPerRank(matrix, rows)).To(Equal(expected))
	})
})
