package nmfsafe

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	goAnnotations    string
	goHierarchy      string
	namespace        string
	nmfLocalization  string
	nmfSummary       string
	outFile          string
	outSummaryFile   string
	safeLocalization string
	safeSummary      string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	goAnnotations := flags.SetString("goAnnotations", args, fileOptions, "")
	goHierarchy := flags.SetString("goHierarchy", args, fileOptions, "")
	namespace := flags.SetString("namespace", args, fileOptions, "CC")
	nmfLocalization := flags.SetString("nmfLocalization", args, fileOptions, "")
	nmfSummary := flags.SetString("nmfSummary", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "concordance.txt")
	outSummaryFile := flags.SetString("outSummaryFile", args, fileOptions, "summary.txt")
	safeLocalization := flags.SetString("safeLocalization", args, fileOptions, "")
	safeSummary := flags.SetString("safeSummary", args, fileOptions, "")

	// Copy arguments from options file.
	options := parameters{
		goAnnotations:    goAnnotations,
		goHierarchy:      goHierarchy,
		namespace:        namespace,
		nmfLocalization:  nmfLocalization,
		nmfSummary:       nmfSummary,
		outFile:          outFile,
		outSummaryFile:   outSummaryFile,
		safeLocalization: safeLocalization,
		safeSummary:      safeSummary,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
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
	if options.safeLocalization == "" {
		messages = append(messages, "missing SAFE localization file")
	}
	if options.safeSummary == "" {
		messages = append(messages, "missing SAFE rank summary file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
