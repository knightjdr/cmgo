package localize

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	database      string
	fdr           float64
	goAnnotations string
	goHierarchy   string
	minBaits      int
	namespace     string
	outFile       string
	saintFile     string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	database := flags.SetString("database", args, fileOptions, "")
	fdr := flags.SetFloat("fdr", args, fileOptions, 0.01)
	goAnnotations := flags.SetString("goAnnotations", args, fileOptions, "")
	goHierarchy := flags.SetString("goHierarchy", args, fileOptions, "")
	minBaits := flags.SetInt("minBaits", args, fileOptions, 1)
	namespace := flags.SetString("namespace", args, fileOptions, "CC")
	outFile := flags.SetString("outFile", args, fileOptions, "lba-localization.txt")
	saintFile := flags.SetString("saintFile", args, fileOptions, "")

	options := parameters{
		database:      database,
		fdr:           fdr,
		goAnnotations: goAnnotations,
		goHierarchy:   goHierarchy,
		namespace:     namespace,
		minBaits:      minBaits,
		outFile:       outFile,
		saintFile:     saintFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.database == "" {
		messages = append(messages, "missing database file")
	}
	if options.goAnnotations == "" {
		messages = append(messages, "missing GO annotations (.gaf) file")
	}
	if options.goHierarchy == "" {
		messages = append(messages, "missing GO hierarchy (.obo) file")
	}
	if options.saintFile == "" {
		messages = append(messages, "missing SAINT file")
	}

	// Format error message.
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
