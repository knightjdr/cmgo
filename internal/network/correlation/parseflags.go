package correlation

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	colorList         string
	cutoff            float64
	edgesPerNode      int
	localizations     string
	maxEdges          int
	nodeLocalizations string
	nodeProfiles      string
	outFile           string
	outFileNetwork    string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	colorList := flags.SetString("colorList", args, fileOptions, "")
	cutoff := flags.SetFloat("cutoff", args, fileOptions, 0)
	edgesPerNode := flags.SetInt("edgesPerNode", args, fileOptions, 20)
	localizations := flags.SetString("localizations", args, fileOptions, "")
	maxEdges := flags.SetInt("maxEdges", args, fileOptions, 0)
	nodeLocalizations := flags.SetString("nodeLocalizations", args, fileOptions, "")
	nodeProfiles := flags.SetString("nodeProfiles", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "corr.txt")
	outFileNetwork := flags.SetString("outFileNetwork", args, fileOptions, "corr.cyjs")

	options := parameters{
		colorList:         colorList,
		cutoff:            cutoff,
		edgesPerNode:      edgesPerNode,
		localizations:     localizations,
		maxEdges:          maxEdges,
		nodeLocalizations: nodeLocalizations,
		nodeProfiles:      nodeProfiles,
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
	if options.nodeLocalizations == "" {
		messages = append(messages, "missing node localization file")
	}
	if options.nodeProfiles == "" {
		messages = append(messages, "missing node profile file")
	}

	// Format error message.
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
