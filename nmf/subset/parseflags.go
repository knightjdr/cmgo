package subset

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/flags"
)

type parameters struct {
	abundanceCap     float64
	basisMatrix      string
	clusteringMethod string
	distanceMetric   string
	minAbundance     float64
	outFile          string
	ranks1           []string
	ranks2           []string
	threshold        float64
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	abundanceCap := flags.SetFloat("abundanceCap", args, fileOptions, 10)
	basisMatrix := flags.SetString("basisMatrix", args, fileOptions, "")
	clusteringMethod := flags.SetString("clusteringMethod", args, fileOptions, "complete")
	distanceMetric := flags.SetString("distanceMetric", args, fileOptions, "euclidean")
	minAbundance := flags.SetFloat("minAbundance", args, fileOptions, 0)
	outFile := flags.SetString("outFile", args, fileOptions, "basis-subset.svg")
	ranks1 := strings.Split(flags.SetString("ranks1", args, fileOptions, ""), ",")
	ranks2 := strings.Split(flags.SetString("ranks2", args, fileOptions, ""), ",")
	threshold := flags.SetFloat("threshold", args, fileOptions, 0.5)

	// Copy arguments from options file.
	options := parameters{
		abundanceCap:     abundanceCap,
		basisMatrix:      basisMatrix,
		clusteringMethod: clusteringMethod,
		distanceMetric:   distanceMetric,
		minAbundance:     minAbundance,
		outFile:          outFile,
		ranks1:           ranks1,
		ranks2:           ranks2,
		threshold:        threshold,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.basisMatrix == "" {
		messages = append(messages, "missing basis matrix file")
	}
	if len(options.ranks1) < 1 || (len(options.ranks1) == 1 && options.ranks1[0] == "") {
		messages = append(messages, "missing first compartment to check")
	}
	if len(options.ranks2) < 1 || (len(options.ranks2) == 1 && options.ranks2[0] == "") {
		messages = append(messages, "missing second compartment to check")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}