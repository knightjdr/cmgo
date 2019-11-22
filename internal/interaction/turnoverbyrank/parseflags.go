package turnoverbyrank

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	fdr          float64
	outFile      string
	saint        string
	species      string
	turnoverFile string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	fdr := flags.SetFloat("fdr", args, fileOptions, 0.01)
	outFile := flags.SetString("outFile", args, fileOptions, "known-by-rank.txt")
	saint := flags.SetString("saint", args, fileOptions, "")
	species := flags.SetString("species", args, fileOptions, "9606")
	turnoverFile := flags.SetString("turnoverFile", args, fileOptions, "")

	// Copy arguments from options file.
	options := parameters{
		fdr:          fdr,
		outFile:      outFile,
		saint:        saint,
		species:      species,
		turnoverFile: turnoverFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.saint == "" {
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
