package goenrich

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	baits     []string
	namespace string
	outFile   string
	saint     string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	baits := flags.SetString("baits", args, fileOptions, "")
	namespace := flags.SetString("namespace", args, fileOptions, "GO:CC")
	outFile := flags.SetString("outFile", args, fileOptions, "go-enrichment.txt")
	saint := flags.SetString("saint", args, fileOptions, "")

	// Copy arguments from options file.
	options := parameters{
		baits:     strings.Split(baits, ","),
		namespace: namespace,
		outFile:   outFile,
		saint:     saint,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if len(options.baits) == 0 {
		messages = append(messages, "missing list of baits")
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
