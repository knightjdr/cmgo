package stats_test

import (
	"github.com/knightjdr/cmgo/pkg/stats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Occurence", func() {
	It("should return a map with the number of times each key occurs in the input slice", func() {
		slice := []string{"a", "a", "b", "c", "c", "d", "c", "d", "d", "d"}
		expected := map[string]int{
			"a": 2,
			"b": 1,
			"c": 3,
			"d": 4,
		}
		Expect(stats.Occurrence(slice)).To(Equal(expected))
	})
})
