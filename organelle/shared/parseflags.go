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
	compartmentFile := flags.ConvertString(args["compartmentFile"])
	fdr := flags.ConvertFloat(args["fdr"])
	minPreyOccurrence := flags.ConvertInt(args["minPreyOccurrence"])
	outFile := flags.ConvertString(args["outFile"])
	regionFile := flags.ConvertString(args["regionFile"])
	saintFile := flags.ConvertString(args["saintFile"])

	// Copy arguments from options file.
	options := parameters{}
	if fileOptions["compartmentFile"] != nil {
		options.compartmentFile = fileOptions["compartmentFile"].(string)
	}
	if fileOptions["fdr"] != nil {
		options.fdr = fileOptions["fdr"].(float64)
	}
	if fileOptions["minPreyOccurrence"] != nil {
		options.minPreyOccurrence = int(fileOptions["minPreyOccurrence"].(float64))
	}
	if fileOptions["outFile"] != nil {
		options.outFile = fileOptions["outFile"].(string)
	}
	if fileOptions["regionFile"] != nil {
		options.regionFile = fileOptions["regionFile"].(string)
	}
	if fileOptions["saintFile"] != nil {
		options.saintFile = fileOptions["saintFile"].(string)
	}

	// Overwrite options file arguments if specified
	if compartmentFile != "" {
		options.compartmentFile = compartmentFile
	}
	if fdr != 0 {
		options.fdr = fdr
	}
	if minPreyOccurrence != 0 {
		options.minPreyOccurrence = minPreyOccurrence
	}
	if outFile != "" {
		options.outFile = outFile
	}
	if regionFile != "" {
		options.regionFile = regionFile
	}
	if saintFile != "" {
		options.saintFile = saintFile
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
	if options.minPreyOccurrence == 0 {
		options.minPreyOccurrence = 1
	}
	if options.outFile == "" {
		options.outFile = "organelle-shared.txt"
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
