package countgo

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
)

func countGenesWithAnnotation(genes []string, annotations map[string]map[string]*geneontology.GOannotation) (int, []string, []string) {
	sum := 0

	withAnnotation := make([]string, 0)
	withoutAnnotation := make([]string, 0)
	for _, gene := range genes {
		if _, ok := annotations[gene]; ok {
			withAnnotation = append(withAnnotation, gene)
			sum++
		} else {
			withoutAnnotation = append(withoutAnnotation, gene)
		}
	}

	return sum, withAnnotation, withoutAnnotation
}
