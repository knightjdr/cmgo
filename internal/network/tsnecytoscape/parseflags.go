package tsnecytoscape

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	colorList         string
	localizations     string
	nodeCoordinates   string
	nodeLocalizations string
	outFile           string
	width             float64
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	colorList := flags.SetString("colorList", args, fileOptions, "")
	localizations := flags.SetString("localizations", args, fileOptions, "")
	nodeCoordinates := flags.SetString("nodeCoordinates", args, fileOptions, "")
	nodeLocalizations := flags.SetString("nodeLocalizations", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "map.cyjs")
	width := flags.SetFloat("width", args, fileOptions, 1000)

	// Copy arguments from options file.
	options := parameters{
		colorList:         colorList,
		localizations:     localizations,
		nodeCoordinates:   nodeCoordinates,
		nodeLocalizations: nodeLocalizations,
		outFile:           outFile,
		width:             width,
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
