package transmembrane

import (
	"github.com/knightjdr/cmgo/pkg/gene"
	"github.com/knightjdr/cmgo/pkg/mapfunc"
	"github.com/knightjdr/cmgo/pkg/pfam"
	"github.com/knightjdr/cmgo/pkg/slice"
	"github.com/knightjdr/cmgo/pkg/uniprot"
)

type orientationData struct {
	Cytosolic int
	Length    int
	Lumenal   int
	UniProt   string
}

func getTransmembraneProteins(genes []string) ([]string, map[string]orientationData) {
	symbolToUniProt := gene.MapIDs(genes, "Symbol", "UniProt", "")
	uniprotIDs := mapfunc.ValuesStringString(symbolToUniProt)

	regions := pfam.GetRegions(uniprotIDs, "")
	transmembraneProteins := filterTransmembrane(regions)
	transmembraneProteinData := uniprot.GetProteins(transmembraneProteins, "")
	orientation := extractOrientation(transmembraneProteinData)

	return mapProteinToGene(symbolToUniProt, orientation)
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

func extractOrientation(entries *uniprot.Entries) map[string]orientationData {
	orientiation := make(map[string]orientationData, 0)

	for id, entry := range *entries {
		orientiation[id] = orientationData{
			Cytosolic: sumFeature(entry.Features, []string{"Cytoplasmic"}),
			Length:    entry.Sequence.Length,
			Lumenal:   sumFeature(entry.Features, []string{"Extracellular", "Lumenal", "Perinuclear space"}),
			UniProt:   id,
		}
	}

	return orientiation
}

func sumFeature(features []uniprot.Feature, description []string) int {
	length := 0

	for _, feature := range features {
		if slice.ContainsString(feature.Description, description) {
			length += feature.End - feature.Begin + 1
		}
	}

	return length
}

func mapProteinToGene(symbolToUniProt map[string]string, orientation map[string]orientationData) ([]string, map[string]orientationData) {
	geneData := make(map[string]orientationData, 0)
	genes := make([]string, 0)

	uniprotToSymbol := mapfunc.ReverseStringString(symbolToUniProt)
	for uniprotID, data := range orientation {
		symbol := uniprotToSymbol[uniprotID]
		geneData[symbol] = data
		genes = append(genes, symbol)
	}

	return genes, geneData
}
