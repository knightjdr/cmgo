package moonlighting

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate moonlighting between ranks", func() {
	It("should return a matrix of primary (row) vs secondary (column) moonlighting counts", func() {
		minRankValue := 0.15
		moonlightingScores := moonScores{
			&preyInfo{PrimaryRank: 1, SecondaryRank: 2, SecondaryScore: 0.18},
			&preyInfo{PrimaryRank: 1, SecondaryRank: 2, SecondaryScore: 0.14},
			&preyInfo{PrimaryRank: 1, SecondaryRank: 2, SecondaryScore: 0.15},
			&preyInfo{PrimaryRank: 0, SecondaryRank: 1, SecondaryScore: 0.25},
			&preyInfo{PrimaryRank: 4, SecondaryRank: 0, SecondaryScore: 0.35},
			&preyInfo{PrimaryRank: 3, SecondaryRank: -1, SecondaryScore: 0.25},
			&preyInfo{PrimaryRank: 3, SecondaryRank: 2, SecondaryScore: 0.25},
		}
		numberOfRanks := 5

		expected := [][]int{
			{0, 1, 0, 0, 0},
			{0, 0, 2, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{1, 0, 0, 0, 0},
		}
		Expect(countRankMoonlighting(moonlightingScores, numberOfRanks, minRankValue)).To(Equal(expected))
	})
})
