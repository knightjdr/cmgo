package preys

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	bait              string
	enrichmentLimit   int
	inter             string
	outFile           string
	outFileEnrichment string
	prey              string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	bait := flags.SetString("bait", args, fileOptions, "")
	enrichmentLimit := flags.SetInt("enrichmentLimit", args, fileOptions, 200)
	inter := flags.SetString("inter", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "control-preys.txt")
	outFileEnrichment := flags.SetString("outFileEnrichment", args, fileOptions, "enrichment.txt")
	prey := flags.SetString("prey", args, fileOptions, "")

	options := parameters{
		bait:              bait,
		enrichmentLimit:   enrichmentLimit,
		inter:             inter,
		outFile:           outFile,
		outFileEnrichment: outFileEnrichment,
		prey:              prey,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.bait == "" {
		messages = append(messages, "missing bait.dat file")
	}
	if options.inter == "" {
		messages = append(messages, "missing inter.dat file")
	}
	if options.prey == "" {
		messages = append(messages, "missing prey.dat file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
