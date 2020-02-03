package rankmetrics

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	"github.com/knightjdr/cmgo/pkg/gene"
	"github.com/knightjdr/cmgo/pkg/slice"
)

func fetchUniprotIDs(saint *saint.SAINT) map[string]string {
	ids := make([]string, 0)
	entrezToSymbol := make(map[string]string, 0)

	for _, row := range *saint {
		entrezToSymbol[row.Prey] = row.PreyGene
		ids = append(ids, row.Prey)
	}
	ids = slice.UniqueStrings(ids)

	entrezToUniprot := gene.MapIDs(ids, "Entrez", "UniProt", "")

	uniprotToGene := make(map[string]string, 0)
	for entrez, uniprot := range entrezToUniprot {
		uniprotToGene[uniprot] = entrezToSymbol[entrez]
	}

	return uniprotToGene
}
