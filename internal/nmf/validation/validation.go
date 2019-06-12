// Package validation assesses the sensitivity of NMF rank assignments
package validation

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/nmf"
)

// Validation performs GO enrichments on each NMF rank and tests how
// sensitive these are to the genes used.
func Validation(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	basis, _, _ := nmf.Basis(options.basisMatrix)
	filterBasis(basis, options.maxGenesPerRank, options.minRankValue, options.withinRankMax)
}
