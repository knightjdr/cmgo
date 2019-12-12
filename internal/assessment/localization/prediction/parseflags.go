package prediction

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	baitExpected          string
	domainsPerCompartment string
	domainsPerGene        string
	fdr                   float64
	fractionation         string
	goHierarchy           string
	hpa                   string
	outFile               string
	predictions           string
	predictionSummary     string
	predictionType        string
	saint                 string
	uniprot               string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	baitExpected := flags.SetString("baitExpected", args, fileOptions, "")
	domainsPerCompartment := flags.SetString("domainsPerCompartment", args, fileOptions, "")
	domainsPerGene := flags.SetString("domainsPerGene", args, fileOptions, "")
	fdr := flags.SetFloat("fdr", args, fileOptions, 0.01)
	fractionation := flags.SetString("fractionation", args, fileOptions, "")
	goHierarchy := flags.SetString("goHierarchy", args, fileOptions, "")
	hpa := flags.SetString("hpa", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "prediction-score.txt")
	predictions := flags.SetString("predictions", args, fileOptions, "")
	predictionSummary := flags.SetString("predictionSummary", args, fileOptions, "")
	predictionType := flags.SetString("predictionType", args, fileOptions, "nmf")
	saint := flags.SetString("saint", args, fileOptions, "")
	uniprot := flags.SetString("uniprot", args, fileOptions, "")

	// Copy arguments from options file.
	options := parameters{
		baitExpected:          baitExpected,
		domainsPerCompartment: domainsPerCompartment,
		domainsPerGene:        domainsPerGene,
		fdr:                   fdr,
		fractionation:         fractionation,
		goHierarchy:           goHierarchy,
		hpa:                   hpa,
		outFile:               outFile,
		predictions:           predictions,
		predictionSummary:     predictionSummary,
		predictionType:        predictionType,
		saint:                 saint,
		uniprot:               uniprot,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.baitExpected == "" {
		messages = append(messages, "missing expected bait localizations")
	}
	if options.domainsPerCompartment == "" {
		messages = append(messages, "missing list of domains per compartment")
	}
	if options.domainsPerGene == "" {
		messages = append(messages, "missing list of domains per gene")
	}
	if options.fractionation == "" {
		messages = append(messages, "missing fractionation predictions")
	}
	if options.goHierarchy == "" {
		messages = append(messages, "missing GO .obo file")
	}
	if options.hpa == "" {
		messages = append(messages, "missing HPA predictions")
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
	if options.uniprot == "" {
		messages = append(messages, "missing UniProt file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
