// Package subset filters an NMF basis (prey) matrix to only include preys enriched all specified rank.
package subset

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/knightjdr/cmgo/cluster"
	"github.com/knightjdr/cmgo/fs"
	"github.com/knightjdr/cmgo/image/svg/heatmap"
	"github.com/knightjdr/cmgo/image/svg/legend"
	"github.com/knightjdr/cmgo/strfunc"
	"github.com/spf13/afero"
)

// NMF filters and NMF basis matrix and outputs an SVG of the result.
func NMF(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	basis, columns, rows := readBasis(options.basisMatrix)

	// Define columns that are specifed by rank names and filter matrix
	rank1Indices, rank2Indices, err := defineColumns(columns, options.ranks1, options.ranks2)
	if err != nil {
		log.Fatalln(err)
	}

	// Remove rows that do not have a maximum in one of the desired ranks, then remove
	// rows that do not have desired rank values within threshold.
	basis, rows = filterByRank(basis, rows, rank1Indices, rank2Indices, options.minNMFScore)
	basis, rows = filterByThreshold(basis, rows, rank1Indices, rank2Indices, options.threshold)

	// Clustering.
	basis, columns, rows = cluster.Process(basis, columns, rows, options.distanceMetric, options.clusteringMethod)

	parameters := heatmap.Settings{
		AbundanceCap: options.abundanceCap,
		Filename:     options.outFile,
		FillColor:    "blueBlack",
		InvertColor:  false,
		MinAbundance: options.minAbundance,
	}
	heatmap.Draw(basis, columns, rows, parameters)

	// Legend.
	dir := filepath.Dir(options.outFile)
	outFile := filepath.Base(options.outFile)

	legendTitle := fmt.Sprintf("NMF value - %s", outFile)
	distanceLegend := legend.Gradient("blueBlack", legendTitle, 101, options.minAbundance, options.abundanceCap, false)

	legendFileName := fmt.Sprintf("%s/legend-%s", dir, outFile)
	afero.WriteFile(fs.Instance, legendFileName, []byte(distanceLegend), 0644)

	// Output preys as a txt file.
	fileName := strfunc.BeforeLast(outFile, ".")
	preyFileName := fmt.Sprintf("%s/preys-%s.txt", dir, fileName)
	preys := strings.Join(rows, "\n")
	afero.WriteFile(fs.Instance, preyFileName, []byte(preys), 0644)
}
