package prediction

import (
	"sort"

	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/pkg/slice"
)

type baitInformation struct {
	localizations     map[string][]int
	compartmentCounts map[int]int
}

func getBaitLocalizationsAsCompartments(baitExpected string, inputFiles fileContent) baitInformation {
	baitLocalizations := localization.Expected(baitExpected)
	compartmentGoTerms := addCompartmentChildren(inputFiles.predictionSummary, inputFiles.goHierarchy)
	return mapLocalizationToCompartment(baitLocalizations, compartmentGoTerms)
}

func addCompartmentChildren(compartmentSummary localization.Summary, goHierarchy *geneontology.GOhierarchy) map[int][]string {
	compartmentGoTerms := make(map[int][]string, 0)

	for compartment, summary := range compartmentSummary {
		compartmentGoTerms[compartment] = make([]string, 0)
		for _, id := range summary.GOid {
			compartmentGoTerms[compartment] = append(compartmentGoTerms[compartment], id)
			compartmentGoTerms[compartment] = append(compartmentGoTerms[compartment], (*goHierarchy)["CC"][id].Children...)
		}
		compartmentGoTerms[compartment] = slice.UniqueStrings(compartmentGoTerms[compartment])
	}

	return compartmentGoTerms
}

func mapLocalizationToCompartment(baitLocalizations localization.ExpectedLocalizations, compartmentGoTerms map[int][]string) baitInformation {
	baitCompartments := make(map[string][]int, 0)
	compartmentCounts := make(map[int]int, 0)

	for bait, localizations := range baitLocalizations {
		baitCompartments[bait] = make([]int, 0)
		for _, localization := range localizations.GoID {
			for rank, ids := range compartmentGoTerms {
				if slice.ContainsString(localization, ids) {
					baitCompartments[bait] = append(baitCompartments[bait], rank)
					compartmentCounts[rank]++
				}
			}
		}
		sort.Ints(baitCompartments[bait])
	}

	return baitInformation{
		compartmentCounts: compartmentCounts,
		localizations:     baitCompartments,
	}
}
