// Package notsignificant outputs non significant preys
package notsignificant

import (
	"log"

	"github.com/knightjdr/cmgo/read/saint"
)

// List reads a SAINT txt file and outputs a list of non-significant preys.
func List(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	saint := saint.Read(options.saintFile, 1, 0)
	saint = removeSignificant(saint, options.fdr)
	summary := summarize(saint)

	writeSummary(summary, options.outFile)
}
