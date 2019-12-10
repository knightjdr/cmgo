// Package prediction calculates a prediction score for each prey.
package prediction

import (
	"log"
)

type preyScore struct {
	Bait   *preyBaitScore
	Domain *preyDomainScore
}

// Score calculates a prediction score for each prey.
func Score(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	inputFiles := readSharedInputFiles(options)

	scores := preyScore{
		Bait:   calculateBaitComponent(options, inputFiles),
		Domain: calculateDomainComponent(options, inputFiles),
	}

	writeScores(scores, options.outFile)
}
