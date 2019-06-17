package robustness

import (
	"errors"
	"strconv"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	basisMatrix     string
	maxGenesPerRank int
	minRankValue    float64
	outFile         string
	outFileSummary  string
	percentiles     []float64
	persistence     float64
	replicates      int
	withinRankMax   float64
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	basisMatrix := flags.SetString("basisMatrix", args, fileOptions, "")
	maxGenesPerRank := flags.SetInt("maxGenesPerRank", args, fileOptions, 100)
	minRankValue := flags.SetFloat("minRankValue", args, fileOptions, 0.25)
	outFile := flags.SetString("outFile", args, fileOptions, "robustness.txt")
	outFileSummary := flags.SetString("outFileSummary", args, fileOptions, "summary.txt")
	percentiles := flags.SetString("percentiles", args, fileOptions, "")
	persistence := flags.SetFloat("persistence", args, fileOptions, 0.9)
	replicates := flags.SetInt("replicates", args, fileOptions, 3)
	withinRankMax := flags.SetFloat("withinRankMax", args, fileOptions, 0.75)

	percentilesSlice := strings.Split(percentiles, ",")
	percentleFloatSlice := make([]float64, len(percentilesSlice))
	for i, value := range percentilesSlice {
		percentleFloatSlice[i], _ = strconv.ParseFloat(value, 64)
	}

	options := parameters{
		basisMatrix:     basisMatrix,
		maxGenesPerRank: maxGenesPerRank,
		minRankValue:    minRankValue,
		outFile:         outFile,
		outFileSummary:  outFileSummary,
		percentiles:     percentleFloatSlice,
		persistence:     persistence,
		replicates:      replicates,
		withinRankMax:   withinRankMax,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.basisMatrix == "" {
		messages = append(messages, "missing basis matrix file")
	}

	// Format error message.
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
