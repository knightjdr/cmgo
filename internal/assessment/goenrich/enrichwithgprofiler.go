package goenrich

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	"github.com/knightjdr/cmgo/pkg/gprofiler"
	"github.com/knightjdr/cmgo/pkg/slice"
)

func enrichWithGprofiler(saint *saint.SAINT, options parameters) map[string][]gprofiler.EnrichedTerm {
	preys := getPreys(options.baits, saint)

	body := gprofiler.RequestBody{
		Organism:      "hsapiens",
		Sources:       []string{options.namespace},
		UserThreshold: 0.01,
	}

	service := gprofiler.Service{
		Body: body,
	}

	enrichment := make(map[string][]gprofiler.EnrichedTerm, 0)
	for _, bait := range options.baits {
		service.Body.Query = preys[bait]
		gprofiler.Fetch(&service)
		enrichment[bait] = service.Result
	}

	return enrichment
}

func getPreys(baits []string, saint *saint.SAINT) map[string][]string {
	baitLookup := slice.Dict(baits)
	preys := make(map[string][]string, 0)

	for _, bait := range baits {
		preys[bait] = make([]string, 0)
	}

	for _, row := range *saint {
		if _, ok := baitLookup[row.Bait]; ok {
			preys[row.Bait] = append(preys[row.Bait], row.PreyGene)
		}
	}

	return preys
}
