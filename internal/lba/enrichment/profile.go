package enrichment

import (
	"github.com/knightjdr/cmgo/pkg/gprofiler"
)

func profile(preyPartners map[string][]string, refseqMap map[string]map[string]string, background []string) map[string][]gprofiler.EnrichedTerm {
	service := gprofiler.Service{}
	// service.Body.Background = background
	service.Body.Ordered = true
	service.Body.Organism = "hsapiens"
	service.Body.Sources = []string{"GO:CC"}

	enrichment := make(map[string][]gprofiler.EnrichedTerm, len(preyPartners))
	for prey, partners := range preyPartners {
		// Convert partner Refseq IDs to Entrez.
		service.Body.Query = make([]string, 0)
		for _, queryRefseqID := range partners {
			if refseqMap[queryRefseqID]["Entrez"] != "" {
				service.Body.Query = append(service.Body.Query, refseqMap[queryRefseqID]["Entrez"])
			}
		}

		gprofiler.Fetch(&service)
		enrichment[prey] = service.Result
	}

	return enrichment
}
