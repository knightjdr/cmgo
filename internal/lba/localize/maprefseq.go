package localize

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/database"
)

func mapRefseq(databaseData []database.Fasta, preys []string) (map[string]map[string]string, []string) {
	mapping := make(map[string]map[string]string, 0)
	for _, prey := range preys {
		mapping[prey] = make(map[string]string, 0)
	}

	geneIDs := make([]string, 0)
	for _, entry := range databaseData {
		if _, ok := mapping[entry.Refseq]; ok {
			mapping[entry.Refseq]["Entrez"] = entry.Entrez
			mapping[entry.Refseq]["Symbol"] = entry.Symbol
			mapping[entry.Refseq]["UniProt"] = ""
			geneIDs = append(geneIDs, entry.Entrez)
		}
	}
	return mapping, geneIDs
}
