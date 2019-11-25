package moonlighting

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	basisMatrix         string
	dissimilarityMatrix string
	minRankValue        float64
	nmfSummary          string
	outFileHeatmap      string
	outFileMatrix       string
	outFileScores       string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	basisMatrix := flags.SetString("basisMatrix", args, fileOptions, "")
	dissimilarityMatrix := flags.SetString("dissimilarityMatrix", args, fileOptions, "")
	minRankValue := flags.SetFloat("minRankValue", args, fileOptions, 0.15)
	nmfSummary := flags.SetString("nmfSummary", args, fileOptions, "")
	outFileHeatmap := flags.SetString("outFileHeatmap", args, fileOptions, "heatmap.svg")
	outFileMatrix := flags.SetString("outFileMatrix", args, fileOptions, "matrix.txt")
	outFileScores := flags.SetString("outFileScores", args, fileOptions, "moonlighting.txt")

	// Copy arguments from options file.
	options := parameters{
		basisMatrix:         basisMatrix,
		dissimilarityMatrix: dissimilarityMatrix,
		minRankValue:        minRankValue,
		nmfSummary:          nmfSummary,
		outFileHeatmap:      outFileHeatmap,
		outFileMatrix:       outFileMatrix,
		outFileScores:       outFileScores,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.basisMatrix == "" {
		messages = append(messages, "missing NMF basis file")
	}
	if options.dissimilarityMatrix == "" {
		messages = append(messages, "missing dissimilarity file")
	}
	if options.nmfSummary == "" {
		messages = append(messages, "missing NMF rank summary file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
