package recovered

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	compartmentID  string
	genes          string
	goAnnotations  string
	localizationID string
	outFile        string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	compartmentID := flags.SetString("compartmentID", args, fileOptions, "")
	genes := flags.SetString("genes", args, fileOptions, "")
	goAnnotations := flags.SetString("goAnnotations", args, fileOptions, "")
	localizationID := flags.SetString("localizationID", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "compartment-recovered.txt")

	// Copy arguments from options file.
	options := parameters{
		compartmentID:  compartmentID,
		genes:          genes,
		goAnnotations:  goAnnotations,
		localizationID: localizationID,
		outFile:        outFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.compartmentID == "" {
		messages = append(messages, "missing GO compartment ID")
	}
	if options.genes == "" {
		messages = append(messages, "missing gene localization file")
	}
	if options.goAnnotations == "" {
		messages = append(messages, "missing GO annotation file")
	}
	if options.localizationID == "" {
		messages = append(messages, "missing localization ID for gene")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
