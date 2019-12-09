// Package prediction calculates a prediction score for each prey.
package prediction

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

// Score calculates a prediction score for each prey.
func Score(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	goHierarchy := geneontology.OBO(options.goHierarchy)
	goHierarchy.GetChildren("CC")

	readBaitLocalizations(options, goHierarchy)

	saintData := saint.Read(options.baitExpected, options.fdr, 0)
	saintData.ParseInteractors(options.fdr)

	getPredictions(options)
}
