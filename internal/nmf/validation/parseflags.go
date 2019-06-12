package validation

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	basisMatrix     string
	maxGenesPerRank int
	minRankValue    float64
	outFile         string
	withinRankMax   float64
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	basisMatrix := flags.SetString("basisMatrix", args, fileOptions, "")
	maxGenesPerRank := flags.SetInt("maxGenesPerRank", args, fileOptions, 100)
	minRankValue := flags.SetFloat("minRankValue", args, fileOptions, 0.25)
	withinRankMax := flags.SetFloat("withinRankMax", args, fileOptions, 0.75)
	outFile := flags.SetString("outFile", args, fileOptions, "basis-subset.svg")

	// Copy arguments from options file.
	options := parameters{
		basisMatrix:     basisMatrix,
		maxGenesPerRank: maxGenesPerRank,
		minRankValue:    minRankValue,
		outFile:         outFile,
		withinRankMax:   withinRankMax,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.basisMatrix == "" {
		messages = append(messages, "missing basis matrix file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
