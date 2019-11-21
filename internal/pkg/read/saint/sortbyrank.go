package saint

import (
	"sort"

	"github.com/knightjdr/cmgo/pkg/mapfunc"
)

// SortByPreyRank returns a list of preys for each bait, sorted in order
// of the spectral count or normalized spectral count
func (s *SAINT) SortByPreyRank(field string) map[string][]string {
	preysAndCountsPerBait := getPreysAndCountsPerBait(s, field)
	return sortPreys(preysAndCountsPerBait)
}

func getPreysAndCountsPerBait(saint *SAINT, field string) map[string]map[string]float64 {
	preysPerBait := make(map[string]map[string]float64)
	for _, row := range *saint {
		if _, ok := preysPerBait[row.Bait]; !ok {
			preysPerBait[row.Bait] = make(map[string]float64, 0)
		}
		preysPerBait[row.Bait][row.PreyGene] = getPreyValue(row, field)
	}
	return preysPerBait
}

func getPreyValue(row Row, field string) float64 {
	var count float64
	if field == "AvgSpec" {
		count = row.AvgSpec
	} else {
		count = row.NormalizedSpec
	}
	return count
}

func sortPreys(preysAndCountsPerBait map[string]map[string]float64) map[string][]string {
	sortedPreysPerBait := make(map[string][]string)
	for bait, preysWithCounts := range preysAndCountsPerBait {
		preys := mapfunc.KeysStringFloat(preysWithCounts)
		sort.Slice(preys, func(i, j int) bool {
			return preysWithCounts[preys[i]] > preysWithCounts[preys[j]]
		})
		sortedPreysPerBait[bait] = preys
	}
	return sortedPreysPerBait
}
