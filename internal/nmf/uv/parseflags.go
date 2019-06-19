package uv

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	basisMatrix     string
	goAnnotations   string
	goHierarchy     string
	maxGenesPerRank int
	minRankValue    float64
	namespace       string
	nmfLocalization string
	nmfSummary      string
	outFile         string
	withinRankMax   float64
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	basisMatrix := flags.SetString("basisMatrix", args, fileOptions, "")
	goAnnotations := flags.SetString("goAnnotations", args, fileOptions, "")
	goHierarchy := flags.SetString("goHierarchy", args, fileOptions, "")
	maxGenesPerRank := flags.SetInt("maxGenesPerRank", args, fileOptions, 100)
	minRankValue := flags.SetFloat("minRankValue", args, fileOptions, 0.25)
	namespace := flags.SetString("namespace", args, fileOptions, "CC")
	nmfLocalization := flags.SetString("nmfLocalization", args, fileOptions, "")
	nmfSummary := flags.SetString("nmfSummary", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "uv-assessment.txt")
	withinRankMax := flags.SetFloat("withinRankMax", args, fileOptions, 0.75)

	options := parameters{
		basisMatrix:     basisMatrix,
		goAnnotations:   goAnnotations,
		goHierarchy:     goHierarchy,
		maxGenesPerRank: maxGenesPerRank,
		minRankValue:    minRankValue,
		namespace:       namespace,
		nmfLocalization: nmfLocalization,
		nmfSummary:      nmfSummary,
		outFile:         outFile,
		withinRankMax:   withinRankMax,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.basisMatrix == "" {
		messages = append(messages, "missing basis matrix file")
	}
	if options.goAnnotations == "" {
		messages = append(messages, "missing GO annotations (.gaf) file")
	}
	if options.goHierarchy == "" {
		messages = append(messages, "missing GO hierarchy (.obo) file")
	}
	if options.nmfLocalization == "" {
		messages = append(messages, "missing NMF localization file")
	}
	if options.nmfSummary == "" {
		messages = append(messages, "missing NMF rank summary file")
	}

	// Format error message.
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
