package moonlighting

import (
	"github.com/knightjdr/cmgo/pkg/mapfunc"
	"github.com/knightjdr/cmgo/pkg/math"
)

type preyInfo struct {
	MoonlightingScore float64
	PrimaryRank       int
	PrimaryScore      float64
	SecondaryRank     int
	SecondaryScore    float64
}

type moonScores []*preyInfo

func calculateMoonlightingScores(basis [][]float64, compatibleRanks []map[int]bool) moonScores {
	moonlightingScores := initializeMoonlightingScores(len(basis))
	for rowIndex, row := range basis {
		primaryRank, primaryScore := findPrimaryRank(row)
		secondaryRank, secondaryScore := findSecondaryRank(row, compatibleRanks[primaryRank])
		moonlightingScores[rowIndex].MoonlightingScore = math.Round(secondaryScore/primaryScore, 0.001)
		moonlightingScores[rowIndex].PrimaryRank = primaryRank
		moonlightingScores[rowIndex].PrimaryScore = primaryScore
		moonlightingScores[rowIndex].SecondaryRank = secondaryRank
		moonlightingScores[rowIndex].SecondaryScore = secondaryScore
	}
	return moonlightingScores
}

func initializeMoonlightingScores(numberOfPreys int) moonScores {
	moonlightingScores := make(moonScores, numberOfPreys)
	for i := range moonlightingScores {
		moonlightingScores[i] = &preyInfo{}
	}
	return moonlightingScores
}

func findPrimaryRank(row []float64) (int, float64) {
	primaryIndex := 0
	for columnIndex, value := range row {
		if value > row[primaryIndex] {
			primaryIndex = columnIndex
		}
	}
	return primaryIndex, row[primaryIndex]
}

func findSecondaryRank(row []float64, compatibleRanks map[int]bool) (int, float64) {
	secondaryIndex := setStartingSecondaryIndex(compatibleRanks)
	if secondaryIndex == -1 {
		return -1, 0
	}

	for columnIndex := range compatibleRanks {
		if row[columnIndex] > row[secondaryIndex] {
			secondaryIndex = columnIndex
		}
	}
	return secondaryIndex, row[secondaryIndex]
}

func setStartingSecondaryIndex(compatibleRanks map[int]bool) int {
	keys := mapfunc.KeysIntBool(compatibleRanks)
	min, err := math.MinSliceInt(keys)
	if err != nil {
		return -1
	}
	return min
}
