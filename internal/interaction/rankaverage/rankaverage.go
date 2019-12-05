// Package rankaverage calculates the average interaction ranks for supplied prey list
package rankaverage

import (
	"fmt"
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

// CalculateRankAverages calculates the average interaction rank for
// each prey in a supplied list.
func CalculateRankAverages(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	preys := readPreys(options.preys)

	saint := saint.Read(options.saint, 1, 0)
	saint.LengthNormalizeSpectralCounts()
	saint.FilterByFDR(options.fdr)
	sortedPreysPerBait := saint.SortByPreyRank("NormalizedSpec")

	rankAverages, mean, sd := calculateAverages(preys, sortedPreysPerBait)
	fmt.Println("Mean:", mean, "SD:", sd)

	writeAverages(rankAverages, options.outFile)
}
