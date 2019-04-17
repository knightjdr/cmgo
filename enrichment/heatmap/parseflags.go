package heatmap

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/flags"
)

type parameters struct {
	clusteringMethod   string
	compartmentSummary string
	distanceMetric     string
	enrichmentFile     string
	outFile            string
	pValue             float64
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	clusteringMethod := flags.ConvertString(args["clusteringMethod"])
	compartmentSummary := flags.ConvertString(args["compartmentSummary"])
	distanceMetric := flags.ConvertString(args["distanceMetric"])
	enrichmentFile := flags.ConvertString(args["enrichmentFile"])
	outFile := flags.ConvertString(args["outFile"])
	pValue := flags.ConvertFloat(args["pValue"])

	// Copy arguments from options file.
	options := parameters{}
	if fileOptions["clusteringMethod"] != nil {
		options.clusteringMethod = fileOptions["clusteringMethod"].(string)
	}
	if fileOptions["compartmentSummary"] != nil {
		options.compartmentSummary = fileOptions["compartmentSummary"].(string)
	}
	if fileOptions["distanceMetric"] != nil {
		options.distanceMetric = fileOptions["distanceMetric"].(string)
	}
	if fileOptions["enrichmentFile"] != nil {
		options.enrichmentFile = fileOptions["enrichmentFile"].(string)
	}
	if fileOptions["outFile"] != nil {
		options.outFile = fileOptions["outFile"].(string)
	}
	if fileOptions["pValue"] != nil {
		options.pValue = fileOptions["pValue"].(float64)
	}

	// Overwrite options file arguments if specified
	if clusteringMethod != "" {
		options.clusteringMethod = clusteringMethod
	}
	if compartmentSummary != "" {
		options.compartmentSummary = compartmentSummary
	}
	if distanceMetric != "" {
		options.distanceMetric = distanceMetric
	}
	if enrichmentFile != "" {
		options.enrichmentFile = enrichmentFile
	}
	if outFile != "" {
		options.outFile = outFile
	}
	if pValue != 0 {
		options.pValue = pValue
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.compartmentSummary == "" {
		messages = append(messages, "missing compartment summary file")
	}
	if options.enrichmentFile == "" {
		messages = append(messages, "missing enriched region file")
	}
	if options.clusteringMethod == "" {
		options.clusteringMethod = "complete"
	}
	if options.distanceMetric == "" {
		options.distanceMetric = "euclidean"
	}
	if options.outFile == "" {
		options.outFile = "region-heatmap.svg"
	}
	if options.pValue == 0 {
		options.pValue = 0.01
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
