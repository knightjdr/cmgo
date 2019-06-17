package robustness

import (
	"math"

	"github.com/knightjdr/cmgo/pkg/gprofiler"
	"github.com/knightjdr/cmgo/pkg/random"
	"github.com/knightjdr/cmgo/pkg/rbo"
)

func dataPoint(geneIndices []int, geneNames, rankDefinition []string, percentiles []float64, persistence float64, replicates int) [][]float64 {
	data := make([][]float64, len(percentiles))

	service := gprofiler.Service{}
	service.Body.Background = geneNames
	service.Body.Organism = "hsapiens"
	service.Body.Sources = []string{"GO:CC"}

	for i, percentile := range percentiles {
		data[i] = make([]float64, replicates)
		percentleLength := int(math.Round(float64(len(geneIndices)) * percentile))

		for j := 0; j < replicates; j++ {
			randomIndices := random.SliceIntN(geneIndices, percentleLength)
			query := make([]string, percentleLength)
			for k, geneIndex := range randomIndices {
				query[k] = geneNames[geneIndex]
			}
			service.Body.Query = query
			fetch(&service)

			definition := make([]string, len(service.Result))
			for j, term := range service.Result {
				definition[j] = term.ID
			}
			data[i][j] = rbo.RBDext(rankDefinition, definition, persistence, 0)
		}
	}
	return data
}
