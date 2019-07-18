package enrichment

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

func associations(saintData []saint.Row) (map[string][]string, map[string][]string) {
	// Store significant baits for each prey and preys for each bait.
	significantBaitsPerPrey := make(map[string][]string, 0)
	significantPreysPerBait := make(map[string][]string, 0)
	for _, row := range saintData {
		significantBaitsPerPrey[row.Prey] = append(significantBaitsPerPrey[row.Prey], row.Bait)
		significantPreysPerBait[row.Bait] = append(significantPreysPerBait[row.Bait], row.Prey)
	}

	return significantBaitsPerPrey, significantPreysPerBait
}
