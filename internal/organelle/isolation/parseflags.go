package isolation

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	abundanceCap      float64
	basisMatrix       string
	correlationCutoff float64
	nmfLocalization   string
	nmfSummary        string
	outFile           string
	svgFile           string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	abundanceCap := flags.SetFloat("abundanceCap", args, fileOptions, 1000)
	basisMatrix := flags.SetString("basisMatrix", args, fileOptions, "")
	correlationCutoff := flags.SetFloat("correlationCutoff", args, fileOptions, 0.9)
	nmfLocalization := flags.SetString("nmfLocalization", args, fileOptions, "")
	nmfSummary := flags.SetString("nmfSummary", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "organelle-isolation.txt")
	svgFile := flags.SetString("svgFile", args, fileOptions, "organelle-isolation.svg")

	options := parameters{
		abundanceCap:      abundanceCap,
		basisMatrix:       basisMatrix,
		correlationCutoff: correlationCutoff,
		nmfLocalization:   nmfLocalization,
		nmfSummary:        nmfSummary,
		outFile:           outFile,
		svgFile:           svgFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.basisMatrix == "" {
		messages = append(messages, "missing basis matrix")
	}
	if options.nmfLocalization == "" {
		messages = append(messages, "missing NMF localization file")
	}
	if options.nmfSummary == "" {
		messages = append(messages, "missing NMF summary file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
