package saint

import (
	"github.com/knightjdr/cmgo/pkg/gene"
	"github.com/knightjdr/cmgo/pkg/slice"
)

// CreateIDmap creates a map of requested accession type to gene names
// using Refseq "Prey" column for the mapping.
func (s *SAINT) CreateIDmap(accessionType string) map[string]string {
	ids := make([]string, 0)
	refseqToSymbol := make(map[string]string, 0)

	for _, row := range *s {
		refseqToSymbol[row.Prey] = row.PreyGene
		ids = append(ids, row.Prey)
	}
	ids = slice.UniqueStrings(ids)

	entrezToUniprot := gene.MapIDs(ids, "Refseq", accessionType, "")

	uniprotToGene := make(map[string]string, 0)
	for refseq, uniprot := range entrezToUniprot {
		uniprotToGene[uniprot] = refseqToSymbol[refseq]
	}

	return uniprotToGene
}
