package recovered

import "github.com/knightjdr/cmgo/internal/pkg/read/geneontology"

func readCompartmentGenes(annotations map[string]map[string]*geneontology.GOannotation, compartmentID string) []string {
	genes := make([]string, 0)

	for gene, terms := range annotations {
		if _, ok := terms[compartmentID]; ok {
			genes = append(genes, gene)
		}
	}

	return genes
}
