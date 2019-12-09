package prediction

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	baitExpected      string
	fdr               float64
	goHierarchy       string
	outFile           string
	predictions       string
	predictionSummary string
	predictionType    string
	saint             string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	baitExpected := flags.SetString("baitExpected", args, fileOptions, "")
	fdr := flags.SetFloat("fdr", args, fileOptions, 0.01)
	goHierarchy := flags.SetString("goHierarchy", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "prediction-score.txt")
	predictions := flags.SetString("predictions", args, fileOptions, "")
	predictionSummary := flags.SetString("predictionSummary", args, fileOptions, "")
	predictionType := flags.SetString("predictionType", args, fileOptions, "nmf")
	saint := flags.SetString("saint", args, fileOptions, "")

	// Copy arguments from options file.
	options := parameters{
		baitExpected:      baitExpected,
		fdr:               fdr,
		goHierarchy:       goHierarchy,
		outFile:           outFile,
		predictions:       predictions,
		predictionSummary: predictionSummary,
		predictionType:    predictionType,
		saint:             saint,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.baitExpected == "" {
		messages = append(messages, "missing expected bait localizations")
	}
	if options.goHierarchy == "" {
		messages = append(messages, "missing GO .obo file")
	}
	if options.predictions == "" {
		messages = append(messages, "missing localization predictions")
	}
	if options.predictionSummary == "" {
		messages = append(messages, "missing prediction summary file")
	}
	if options.saint == "" {
		messages = append(messages, "missing SAINT file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
