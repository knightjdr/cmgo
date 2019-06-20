// Package hydropathy calculates the average hydropathy of proteins sequences
package hydropathy

import (
	"fmt"
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/bioplex"
	"github.com/knightjdr/cmgo/internal/pkg/read/database"
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	"github.com/knightjdr/cmgo/pkg/stats"
)

// Assess the hyrdropathy of protein sequences in a SAINT and
// BioPlex file using the scale of Kyte J. and Doolittle R.F.
// J. Mol. Biol. 157:105-132(1982).
func Assess(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	databaseData := database.Read(options.database, true)
	entrezSequenceMap, refseqSequenceMap, refseqEntrez := sequenceMapping(databaseData)

	saintData := saint.Read(options.saintFile, options.fdr, 1)
	preys := uniquePreys(saintData)

	bioplexData := bioplex.Read(options.bioplexFile)
	bioplexInteractors := uniqueBioplex(bioplexData)

	uniqueToSaint := diff(preys, bioplexInteractors, refseqEntrez)

	hydropathySaint := proteinHydropathy(preys, refseqSequenceMap)
	hydropathyBioplex := proteinHydropathy(bioplexInteractors, entrezSequenceMap)
	hydropathyUnique := proteinHydropathy(uniqueToSaint, refseqSequenceMap)
	fmt.Println("Avg. hydropathy SAINT: ", stats.MeanFloat(hydropathySaint))
	fmt.Println("Avg. hydropathy BioPLEX:", stats.MeanFloat(hydropathyBioplex))
	fmt.Println("Avg. hydropathy unique to SAINT:", stats.MeanFloat(hydropathyUnique))
}
