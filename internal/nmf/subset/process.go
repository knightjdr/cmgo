// Package subset filters an NMF basis (prey) matrix to only include preys enriched all specified rank.
package subset

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/knightjdr/cmgo/internal/pkg/cluster"
	"github.com/knightjdr/cmgo/internal/pkg/image/svg/heatmap"
	"github.com/knightjdr/cmgo/internal/pkg/image/svg/legend"
	"github.com/knightjdr/cmgo/internal/pkg/read/nmf"
	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/strfunc"
	"github.com/spf13/afero"
)

// NMF filters and NMF basis matrix and outputs an SVG of the result.
func NMF(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	basis, columns, rows := nmf.ReadBasis(options.basisMatrix)

	// Define columns that are specifed by rank names
	rank1Indices, rank2Indices, err := defineColumns(columns, options.ranks1, options.ranks2)
	if err != nil {
		log.Fatalln(err)
	}

	// Remove rows that do not have a maximum in one of the desired ranks,
	// then rows that do not have desired rank values within threshold, and finally
	// remove rows where the fold-change differential between the desired ranks
	// and the rest is less than the specified cutoff.
	basis, rows = filterByRank(basis, rows, rank1Indices, rank2Indices, options.minNMFScore)
	basis, rows = filterByThreshold(basis, rows, rank1Indices, rank2Indices, options.threshold)
	basis, rows = filterBySpecificity(basis, rows, rank1Indices, rank2Indices, options.specificity)

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
