package heatmap

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/flags"
)

type parameters struct {
	abundanceCap       float64
	clusteringMethod   string
	compartmentSummary string
	distanceMetric     string
	enrichmentFile     string
	minAbundance       float64
	outFile            string
	pValue             float64
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	abundanceCap := flags.SetFloat("abundanceCap", args, fileOptions, 10)
	clusteringMethod := flags.SetString("clusteringMethod", args, fileOptions, "complete")
	compartmentSummary := flags.SetString("compartmentSummary", args, fileOptions, "")
	distanceMetric := flags.SetString("distanceMetric", args, fileOptions, "euclidean")
	enrichmentFile := flags.SetString("enrichmentFile", args, fileOptions, "")
	minAbundance := flags.SetFloat("minAbundance", args, fileOptions, 0)
	outFile := flags.SetString("outFile", args, fileOptions, "region-heatmap.svg")
	pValue := flags.SetFloat("pValue", args, fileOptions, 0.01)

	// Copy arguments from options file.
	options := parameters{
		abundanceCap:       abundanceCap,
		clusteringMethod:   clusteringMethod,
		compartmentSummary: compartmentSummary,
		distanceMetric:     distanceMetric,
		enrichmentFile:     enrichmentFile,
		minAbundance:       minAbundance,
		outFile:            outFile,
		pValue:             pValue,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.compartmentSummary == "" {
		messages = append(messages, "missing compartment summary file")
	}
	if options.enrichmentFile == "" {
		messages = append(messages, "missing enriched region file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
