// Package moonlighting scores preys localized by NMF for moonlighting in more than one compartment
package moonlighting

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/nmf"
	readNMF "github.com/knightjdr/cmgo/internal/pkg/read/nmf"
)

// Calculate calculates prey moonlighting scores from an NMF basis file
// and a dissimilarity matrix.
func Calculate(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	basis, _, rows := readNMF.ReadBasis(options.nmfBasis)
	basis, rows = nmf.FilterBasisByTreshold(basis, rows, options.minimumNmfScore)
}
