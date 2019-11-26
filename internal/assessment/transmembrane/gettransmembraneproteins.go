package transmembrane

import (
	"github.com/knightjdr/cmgo/pkg/gene"
	"github.com/knightjdr/cmgo/pkg/mapfunc"
	"github.com/knightjdr/cmgo/pkg/pfam"
)

func getTransmembraneProteins(genes []string) []string {
	symbolToUniProt := gene.MapIDs(genes, "Symbol", "UniProt", "")
	uniprotIDs := mapfunc.ValuesStringString(symbolToUniProt)
	regions := pfam.GetRegions(uniprotIDs, "")

	transmembraneProteins := filterTransmembrane(regions)

	return mapUniprotToSymbol(symbolToUniProt, transmembraneProteins)
}

func filterTransmembrane(regions *pfam.Features) []string {
	transmembraneProteins := make([]string, 0)

	for uniprotID, features := range *regions {
		for _, motif := range features.Motifs {
			if motif.Name == "transmembrane" {
				transmembraneProteins = append(transmembraneProteins, uniprotID)
				break
			}
		}
	}

	return transmembraneProteins
}

func mapUniprotToSymbol(symbolToUniProt map[string]string, uniprotIDs []string) []string {
	symbols := make([]string, 0)

	uniprotToSymbol := mapfunc.ReverseStringString(symbolToUniProt)
	for _, uniprotID := range uniprotIDs {
		symbols = append(symbols, uniprotToSymbol[uniprotID])
	}

	return symbols
}
