package recovered

import "github.com/knightjdr/cmgo/pkg/slice"

func countRecoveredGenes(localizedGenes, compartmentGenes []string) map[string]bool {
	summary := make(map[string]bool, len(compartmentGenes))

	for _, gene := range compartmentGenes {
		recovered := false
		if slice.ContainsString(gene, localizedGenes) {
			recovered = true
		}
		summary[gene] = recovered
	}

	return summary
}
