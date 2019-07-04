package localize

import (
	customSort "github.com/knightjdr/cmgo/pkg/sort"
)

func topPreyPartners(
	baitsPerPrey map[string][]string,
	preysPerBait map[string][]string,
	preyLimit int,
	minFC float64,
) (map[string]map[string]float64, map[string][]string) {
	foldChange := make(map[string]map[string]float64, 0)
	topPreys := make(map[string][]string, 0)
	for prey, baitList := range baitsPerPrey {
		// Count all partner preys the prey was seen with across the baits
		// it was detected with. Remove the prey itself from map.
		partnerCount := make(map[string]int, 0)
		for _, bait := range baitList {
			for _, partnerPrey := range preysPerBait[bait] {
				partnerCount[partnerPrey]++
			}
		}
		delete(partnerCount, prey)

		// Calculate the fold change for each partner prey.
		numberBaitsForSelectedPrey := len(baitList)
		totalBaits := len(preysPerBait)
		foldChange[prey] = make(map[string]float64, 0)
		for partnerPrey, count := range partnerCount {
			foldChangeBaitSubset := float64(count) / float64(numberBaitsForSelectedPrey)
			foldChangeDataset := float64(len(baitsPerPrey[partnerPrey])) / float64(totalBaits)
			foldChangeForSelectedPrey := foldChangeBaitSubset / foldChangeDataset
			if foldChangeForSelectedPrey > minFC {
				foldChange[prey][partnerPrey] = foldChangeForSelectedPrey
			}
		}

		// Filter the list to only keep the top "preyLimit" entries.
		limit := preyLimit
		if len(foldChange[prey]) < limit {
			limit = len(foldChange[prey])
		}
		topPreys[prey] = make([]string, limit)
		order := customSort.ByMapValueStringFloat(foldChange[prey], "descending")
		for i, entry := range order {
			if i >= limit {
				break
			}
			topPreys[prey][i] = entry.Key
		}
	}

	return foldChange, topPreys
}
