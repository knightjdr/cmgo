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
	for poi, baitList := range baitsPerPrey {
		// Count all partner preys the prey of interest (POI) was seen with across the baits
		// it was detected with. Remove the POI itself from map.
		partnerCount := make(map[string]int, 0)
		for _, bait := range baitList {
			for _, partnerPrey := range preysPerBait[bait] {
				partnerCount[partnerPrey]++
			}
		}
		delete(partnerCount, poi)

		// Calculate the fold change for each partner prey and ignore those with
		// a value below minFC threshold (exclusive).
		foldChange[poi] = make(map[string]float64, 0)
		numberBaitsForPOI := len(baitList)
		totalBaits := len(preysPerBait)
		for partnerPrey, count := range partnerCount {
			foldChangeBaitSubset := float64(count) / float64(numberBaitsForPOI)
			foldChangeDataset := float64(len(baitsPerPrey[partnerPrey])) / float64(totalBaits)
			foldChangeForPartnerPrey := foldChangeBaitSubset / foldChangeDataset
			if foldChangeForPartnerPrey > minFC {
				foldChange[poi][partnerPrey] = foldChangeForPartnerPrey
			} else {
				delete(partnerCount, partnerPrey)
			}
		}

		// Filter the list to only keep the top "preyLimit" entries.
		limit := preyLimit
		if len(partnerCount) < limit {
			limit = len(partnerCount)
		}
		topPreys[poi] = make([]string, limit)
		order := customSort.ByMapValueStringInt(partnerCount, "descending")
		for i, entry := range order {
			if i >= limit {
				break
			}
			topPreys[poi][i] = entry.Key
		}
	}

	return foldChange, topPreys
}
