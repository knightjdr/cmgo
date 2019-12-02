package countgo

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	genes         string
	goAnnotations string
	namespace     string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	genes := flags.SetString("genes", args, fileOptions, "")
	goAnnotations := flags.SetString("goAnnotations", args, fileOptions, "")
	namespace := flags.SetString("namespace", args, fileOptions, "CC")

	// Copy arguments from options file.
	options := parameters{
		genes:         genes,
		goAnnotations: goAnnotations,
		namespace:     namespace,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.genes == "" {
		messages = append(messages, "missing list of genes file")
	}
	if options.goAnnotations == "" {
		messages = append(messages, "missing GO annotation file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
