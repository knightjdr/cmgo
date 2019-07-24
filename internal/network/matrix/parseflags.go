package matrix

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	colorList         string
	cutoff            float64
	localizations     string
	matrix            string
	nodeLocalizations string
	outFile           string
	outFileNetwork    string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	colorList := flags.SetString("colorList", args, fileOptions, "")
	cutoff := flags.SetFloat("cutoff", args, fileOptions, 0.01)
	localizations := flags.SetString("localizations", args, fileOptions, "")
	matrix := flags.SetString("matrix", args, fileOptions, "")
	nodeLocalizations := flags.SetString("nodeLocalizations", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "corr.txt")
	outFileNetwork := flags.SetString("outFileNetwork", args, fileOptions, "corr.cyjs")

	options := parameters{
		colorList:         colorList,
		cutoff:            cutoff,
		localizations:     localizations,
		matrix:            matrix,
		nodeLocalizations: nodeLocalizations,
		outFile:           outFile,
		outFileNetwork:    outFileNetwork,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.colorList == "" {
		messages = append(messages, "missing color list")
	}
	if options.localizations == "" {
		messages = append(messages, "missing localization file")
	}
	if options.matrix == "" {
		messages = append(messages, "missing matrix file")
	}
	if options.nodeLocalizations == "" {
		messages = append(messages, "missing node localization file")
	}

	// Format error message.
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
