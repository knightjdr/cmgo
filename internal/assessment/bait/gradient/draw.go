// Package gradient draws a similarity gradient for baits.
package gradient

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/list"
	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
)

// Draw takes an ordered bait list and expected localizations for the baits
// and creates a gradient showing similarity between adjacent baits.
func Draw(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	baits := list.ParseSlice(options.baitList)
	expectedLocalizations := localization.Expected(options.expectedLocalizations)

	similarity := adjacentSimilarity(baits, expectedLocalizations)

	writeGradient(similarity, options.outFile)
}
