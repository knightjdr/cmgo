package genes

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	background string
	genes      string
	namespace  string
	outFile    string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	background := flags.SetString("background", args, fileOptions, "")
	genes := flags.SetString("genes", args, fileOptions, "")
	namespace := flags.SetString("namespace", args, fileOptions, "CC")
	outFile := flags.SetString("outFile", args, fileOptions, "enrichment.txt")

	// Copy arguments from options file.
	options := parameters{
		background: background,
		genes:      genes,
		namespace:  namespace,
		outFile:    outFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.background == "" {
		messages = append(messages, "missing background file")
	}
	if options.genes == "" {
		messages = append(messages, "missing gene list file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
