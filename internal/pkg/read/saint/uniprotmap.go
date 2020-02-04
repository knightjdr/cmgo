package saint

import (
	"github.com/knightjdr/cmgo/pkg/gene"
	"github.com/knightjdr/cmgo/pkg/slice"
)

var mapRefseqToUniProt = gene.RefseqToUniProt

// GetUniProtMapping creates a map from UniProt accessions to
// gene names using the Refseq "Prey" column IDs.
func (s *SAINT) GetUniProtMapping() map[string]string {
	ids := make([]string, 0)
	refseqToSymbol := make(map[string]string, 0)

	for _, row := range *s {
		refseqToSymbol[row.Prey] = row.PreyGene
		ids = append(ids, row.Prey)
	}
	ids = slice.UniqueStrings(ids)

	refseqToUniprot := mapRefseqToUniProt(ids, "")

	uniprotToGene := make(map[string]string, 0)
	for refseq, accessions := range refseqToUniprot {
		for _, accession := range accessions {
			uniprotToGene[accession] = refseqToSymbol[refseq]
		}
	}

	return uniprotToGene
}
