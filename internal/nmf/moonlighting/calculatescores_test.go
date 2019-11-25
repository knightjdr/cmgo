package moonlighting

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate moonlighting scores", func() {
	It("should return primary and secondary rank information, as well as moonlighting score", func() {
		basis := [][]float64{
			{0, 0.5, 0.25, 0.1, 0.05},
			{1, 0.75, 0.2, 0.8, 0.5},
			{0.1, 0.1, 0.1, 0.1, 0.2},
			{0.1, 0.1, 0.1, 0.5, 0.2},
		}
		compatibleRanks := []map[int]bool{
			{1: true, 2: true, 4: true},
			{0: true, 2: true},
			{0: true, 1: true, 4: true},
			{},
			{0: true, 2: true, 3: true},
		}

		expected := moonScores{
			&preyInfo{
				MoonlightingScore: 0.5,
				PrimaryRank:       1,
				PrimaryScore:      0.5,
				SecondaryRank:     2,
				SecondaryScore:    0.25,
			},
			&preyInfo{
				MoonlightingScore: 0.75,
				PrimaryRank:       0,
				PrimaryScore:      1,
				SecondaryRank:     1,
				SecondaryScore:    0.75,
			},
			&preyInfo{
				MoonlightingScore: 0.5,
				PrimaryRank:       4,
				PrimaryScore:      0.2,
				SecondaryRank:     0,
				SecondaryScore:    0.1,
			},
			&preyInfo{
				MoonlightingScore: 0,
				PrimaryRank:       3,
				PrimaryScore:      0.5,
				SecondaryRank:     -1,
				SecondaryScore:    0,
			},
		}
		Expect(calculateMoonlightingScores(basis, compatibleRanks)).To(Equal(expected))
	})
})
