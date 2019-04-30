package notsignificant

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/flags"
)

type parameters struct {
	fdr       float64
	outFile   string
	saintFile string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	fdr := flags.SetFloat("fdr", args, fileOptions, 0.01)
	outFile := flags.SetString("outFile", args, fileOptions, "not-significant.txt")
	saintFile := flags.SetString("saint", args, fileOptions, "")

	// Copy arguments from options file.
	options := parameters{
		fdr:       fdr,
		outFile:   outFile,
		saintFile: saintFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
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
