package robustness

import (
	"github.com/knightjdr/cmgo/pkg/gprofiler"
)

func defineRanks(characterizingGenes [][]int, geneNames []string) [][]string {
	service := gprofiler.Service{}
	service.Body.Background = geneNames
	service.Body.Organism = "hsapiens"
	service.Body.Sources = []string{"GO:CC"}

	definitions := make([][]string, len(characterizingGenes))
	for i, geneIndices := range characterizingGenes {
		query := make([]string, 0)
		for _, index := range geneIndices {
			query = append(query, geneNames[index])
		}
		service.Body.Query = query
		fetch(&service)

		definitions[i] = make([]string, len(service.Result))
		for j, term := range service.Result {
			definitions[i][j] = term.ID
		}
	}

	return definitions
}
