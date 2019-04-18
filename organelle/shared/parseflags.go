package shared

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/flags"
)

type parameters struct {
	compartmentFile   string
	fdr               float64
	minPreyOccurrence int
	outFile           string
	regionFile        string
	saintFile         string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	compartmentFile := flags.SetString("compartmentFile", args, fileOptions, "")
	fdr := flags.SetFloat("fdr", args, fileOptions, 0.01)
	minPreyOccurrence := flags.SetInt("minPreyOccurrence", args, fileOptions, 1)
	outFile := flags.SetString("outFile", args, fileOptions, "organelle-shared.txt")
	regionFile := flags.SetString("regionFile", args, fileOptions, "")
	saintFile := flags.SetString("saintFile", args, fileOptions, "")

	// Copy arguments from options file.
	options := parameters{
		compartmentFile:   compartmentFile,
		fdr:               fdr,
		minPreyOccurrence: minPreyOccurrence,
		outFile:           outFile,
		regionFile:        regionFile,
		saintFile:         saintFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.compartmentFile == "" {
		messages = append(messages, "missing JSON file with list of compartments to compare")
	}
	if options.regionFile == "" {
		messages = append(messages, "missing region file")
	}
	if options.saintFile == "" {
		messages = append(messages, "missing SAINT file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
