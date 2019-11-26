package transmembrane

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	basisMatrix           string
	cytosolicBaits        []string
	cytosolicCompartments []string
	fdr                   float64
	lumenalBaits          []string
	lumenalCompartments   []string
	minRankValue          float64
	outFile               string
	saint                 string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	basisMatrix := flags.SetString("basisMatrix", args, fileOptions, "")
	cytosolicBaits := strings.Split(flags.SetString("cytosolicBaits", args, fileOptions, ""), ",")
	cytosolicCompartments := strings.Split(flags.SetString("cytosolicCompartments", args, fileOptions, ""), ",")
	fdr := flags.SetFloat("fdr", args, fileOptions, 0.01)
	lumenalBaits := strings.Split(flags.SetString("lumenalBaits", args, fileOptions, ""), ",")
	lumenalCompartments := strings.Split(flags.SetString("lumenalCompartments", args, fileOptions, ""), ",")
	minRankValue := flags.SetFloat("minRankValue", args, fileOptions, 0.15)
	outFile := flags.SetString("outFile", args, fileOptions, "transmembrane.txt")
	saint := flags.SetString("saint", args, fileOptions, "")

	// Copy arguments from options file.
	options := parameters{
		basisMatrix:           basisMatrix,
		cytosolicBaits:        cytosolicBaits,
		cytosolicCompartments: cytosolicCompartments,
		fdr:                   fdr,
		lumenalBaits:          lumenalBaits,
		lumenalCompartments:   lumenalCompartments,
		minRankValue:          minRankValue,
		outFile:               outFile,
		saint:                 saint,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.basisMatrix == "" {
		messages = append(messages, "missing NMF basis file")
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
