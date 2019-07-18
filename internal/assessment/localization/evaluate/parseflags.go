package evaluate

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	goAnnotations  string
	goHierarchy    string
	localization   string
	namespace      string
	outFile        string
	outFileSummary string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	goAnnotations := flags.SetString("goAnnotations", args, fileOptions, "")
	goHierarchy := flags.SetString("goHierarchy", args, fileOptions, "")
	localization := flags.SetString("localization", args, fileOptions, "")
	namespace := flags.SetString("namespace", args, fileOptions, "CC")
	outFile := flags.SetString("outFile", args, fileOptions, "localization-known.txt")
	outFileSummary := flags.SetString("outFileSummary", args, fileOptions, "summary.txt")

	options := parameters{
		goAnnotations:  goAnnotations,
		goHierarchy:    goHierarchy,
		localization:   localization,
		namespace:      namespace,
		outFile:        outFile,
		outFileSummary: outFileSummary,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.goAnnotations == "" {
		messages = append(messages, "missing GO annotations (.gaf) file")
	}
	if options.goHierarchy == "" {
		messages = append(messages, "missing GO hierarchy (.obo) file")
	}
	if options.localization == "" {
		messages = append(messages, "missing localization file")
	}

	// Format error message.
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
