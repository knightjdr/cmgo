package localize

import (
	"sort"

	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unique preys", func() {
	It("should return a slice of unique preys from SAINT data", func() {
		saintData := []saint.Row{
			{Bait: "a", Prey: "W"},
			{Bait: "a", Prey: "X"},
			{Bait: "a", Prey: "Y"},
			{Bait: "b", Prey: "W"},
			{Bait: "b", Prey: "X"},
			{Bait: "b", Prey: "Z"},
		}
		actual := uniquePreys(saintData)
		sort.Strings(actual)
		expected := []string{"W", "X", "Y", "Z"}
		Expect(actual).To(Equal(expected))
	})
})
