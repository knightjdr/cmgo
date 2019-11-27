// Package transmembrane reports the orientation evidence for transmembrane preys.
package transmembrane

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/nmf"
	readNMF "github.com/knightjdr/cmgo/internal/pkg/read/nmf"
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

// Orientation takes all transmembrane proteins residing between a pair/set of organelles
// and reports evidence for cytosolic or lumenal orientation/regions.
func Orientation(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	saintData := saint.Read(options.saint, options.fdr, 0)

	basis, _, preys := readNMF.ReadBasis(options.basisMatrix)
	basis, preys = nmf.FilterBasisByTreshold(basis, preys, options.minRankValue)

	preysPerRank := nmf.GetPreysPerRank(basis, preys)
	cytosolicPreys := getPreysInCompartment(preysPerRank, options.cytosolicCompartments)
	lumenalPreys := getPreysInCompartment(preysPerRank, options.lumenalCompartments)
	transmembranePreys, transmembranePreyData := getTransmembraneProteins(append(cytosolicPreys, lumenalPreys...))

	baitsPerPrey := findBaitsPerPrey(transmembranePreys, saintData)
	organelleBaitsPerPrey := countOrganelleBaitsPerPrey(baitsPerPrey, options.cytosolicBaits, options.lumenalBaits)

	sumOptions := summaryOptions{
		basis:                 basis,
		cytosolicCompartments: options.cytosolicCompartments,
		cytosolicPreys:        cytosolicPreys,
		lumenalCompartments:   options.lumenalCompartments,
		lumenalPreys:          lumenalPreys,
		organelleBaitsPerPrey: organelleBaitsPerPrey,
		rows:                  preys,
		transmembranePreys:    transmembranePreys,
		transmembranePreyData: transmembranePreyData,
	}
	summary := summarize(sumOptions)
	writeSummary(summary, options.outFile)
}
