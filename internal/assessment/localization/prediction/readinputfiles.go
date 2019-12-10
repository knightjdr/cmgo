package prediction

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	"github.com/knightjdr/cmgo/internal/pkg/read/uniprot"
)

type fileContent struct {
	baitInteractors   map[string][]string
	geneToUniProt     map[string]string
	goHierarchy       *geneontology.GOhierarchy
	predictions       map[string]int
	predictionSummary localization.Summary
}

func readSharedInputFiles(options parameters) fileContent {
	goHierarchy := geneontology.OBO(options.goHierarchy)
	goHierarchy.GetChildren("CC")

	predictions := getPredictions(options)

	predictionSummary := localization.SummaryFile(options.predictionSummary)

	saintData := saint.Read(options.saint, options.fdr, 0)
	baitInteractors := saintData.ParseInteractors(options.fdr)

	uniprotData := uniprot.Read(options.uniprot, 9606)

	return fileContent{
		baitInteractors:   baitInteractors,
		geneToUniProt:     uniprotData.CreateGeneNameMap(),
		goHierarchy:       goHierarchy,
		predictions:       predictions,
		predictionSummary: predictionSummary,
	}
}
