package localize

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

func countAssociations(saintData []saint.Row) map[string]map[string]int {
	// Store significant preys for each bait.
	significantPreys := make(map[string][]string, 0)
	for _, row := range saintData {
		significantPreys[row.Bait] = append(significantPreys[row.Bait], row.Prey)
	}

	// Count prey associations.
	preyCounts := make(map[string]map[string]int, 0)
	for _, preys := range significantPreys {
		for i, prey := range preys {
			if _, ok := preyCounts[prey]; !ok {
				preyCounts[prey] = make(map[string]int, 0)
			}
			for j, associatedPrey := range preys {
				if j != i {
					preyCounts[prey][associatedPrey]++
				}
			}
		}
	}

	return preyCounts
}
