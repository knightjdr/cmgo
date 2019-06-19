package uv

import (
	"github.com/knightjdr/cmgo/internal/pkg/localization"
	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	readLocalization "github.com/knightjdr/cmgo/internal/pkg/read/localization"
)

func assessGenes(
	nonCharacterizingGenes [][]string,
	nmfSummary readLocalization.Summary,
	goAnnotations map[string]map[string]*geneontology.GOannotation,
	goHierarchy map[string]*geneontology.GOterm,
) []map[string][]string {
	assessment := make([]map[string][]string, len(nonCharacterizingGenes))
	for i, genes := range nonCharacterizingGenes {
		rank := i + 1
		assessment[i] = map[string][]string{
			"known":   make([]string, 0),
			"unknown": make([]string, 0),
		}
		for _, gene := range genes {
			known := localization.IsKnown(gene, nmfSummary[rank].GOid, goAnnotations, goHierarchy)
			if known {
				assessment[i]["known"] = append(assessment[i]["known"], gene)
			} else {
				assessment[i]["unknown"] = append(assessment[i]["unknown"], gene)
			}
		}
	}
	return assessment
}
