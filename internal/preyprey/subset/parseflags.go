package subset

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	genes   string
	heatmap string
	outFile string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	genes := flags.SetString("genes", args, fileOptions, "")
	heatmap := flags.SetString("heatmap", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "cluster.tsv")

	// Copy arguments from options file.
	options := parameters{
		genes:   genes,
		heatmap: heatmap,
		outFile: outFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.genes == "" {
		messages = append(messages, "missing gene list file")
	}
	if options.heatmap == "" {
		messages = append(messages, "missing heatmap file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
