package knownbyrank

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	biogrid string
	fdr     float64
	intact  string
	outFile string
	saint   string
	species string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	biogrid := flags.SetString("biogrid", args, fileOptions, "")
	fdr := flags.SetFloat("fdr", args, fileOptions, 0.01)
	intact := flags.SetString("intact", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "known-by-rank.txt")
	saint := flags.SetString("saint", args, fileOptions, "")
	species := flags.SetString("species", args, fileOptions, "9606")

	// Copy arguments from options file.
	options := parameters{
		biogrid: biogrid,
		fdr:     fdr,
		intact:  intact,
		outFile: outFile,
		saint:   saint,
		species: species,
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
