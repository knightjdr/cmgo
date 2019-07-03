package localize

import (
	"github.com/knightjdr/cmgo/pkg/gprofiler"
)

func profile(preyPartners map[string][]string, background []string) map[string][]gprofiler.EnrichedTerm {
	service := gprofiler.Service{}
	// service.Body.Background = background
	service.Body.Ordered = true
	service.Body.Organism = "hsapiens"
	service.Body.Sources = []string{"GO:CC"}

	testSet := map[string]bool{
		"NP_001136120.1": true, // SAR1A
		"NP_689485.1":    true, // CHMP7
	}

	enrichment := make(map[string][]gprofiler.EnrichedTerm, len(preyPartners))
	for prey, partners := range preyPartners {
		if _, ok := testSet[prey]; ok {
			service.Body.Query = partners
			gprofiler.Fetch(&service)
			enrichment[prey] = service.Result
		}
	}

	return enrichment
}
