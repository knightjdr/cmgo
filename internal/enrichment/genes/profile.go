package genes

import (
	"fmt"

	"github.com/knightjdr/cmgo/pkg/gprofiler"
)

func profile(genes, background []string, namespace string) []gprofiler.EnrichedTerm {
	service := gprofiler.Service{}
	service.Body.Background = background
	service.Body.Organism = "hsapiens"
	service.Body.Query = genes
	service.Body.Sources = []string{fmt.Sprintf("GO:%s", namespace)}

	gprofiler.Fetch(&service)

	return service.Result
}
