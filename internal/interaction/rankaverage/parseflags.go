package rankaverage

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	fdr     float64
	outFile string
	preys   string
	saint   string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	fdr := flags.SetFloat("fdr", args, fileOptions, 0.01)
	outFile := flags.SetString("outFile", args, fileOptions, "prey-rank-average.txt")
	preys := flags.SetString("preys", args, fileOptions, "")
	saint := flags.SetString("saint", args, fileOptions, "")

	// Copy arguments from options file.
	options := parameters{
		fdr:     fdr,
		outFile: outFile,
		preys:   preys,
		saint:   saint,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.preys == "" {
		messages = append(messages, "missing prey list")
	}
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
