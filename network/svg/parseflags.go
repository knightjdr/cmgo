package svg

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/flags"
)

type parameters struct {
	colorList       string
	localizations   string
	nodeCoordinates string
	outFile         string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	colorList := flags.SetString("colorList", args, fileOptions, "")
	localizations := flags.SetString("localizations", args, fileOptions, "")
	nodeCoordinates := flags.SetString("nodeCoordinates", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "network.svg")

	// Copy arguments from options file.
	options := parameters{
		colorList:       colorList,
		localizations:   localizations,
		nodeCoordinates: nodeCoordinates,
		outFile:         outFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.colorList == "" {
		messages = append(messages, "missing color list")
	}
	if options.localizations == "" {
		messages = append(messages, "missing localization file")
	}
	if options.nodeCoordinates == "" {
		messages = append(messages, "missing node coordinates")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
