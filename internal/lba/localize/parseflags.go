package localize

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	database  string
	fdr       float64
	minBaits  int
	minFC     float64
	namespace string
	outFile   string
	saintFile string
	preyLimit int
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	database := flags.SetString("database", args, fileOptions, "")
	fdr := flags.SetFloat("fdr", args, fileOptions, 0.01)
	minBaits := flags.SetInt("minBaits", args, fileOptions, 1)
	minFC := flags.SetFloat("minFC", args, fileOptions, 1)
	namespace := flags.SetString("namespace", args, fileOptions, "CC")
	outFile := flags.SetString("outFile", args, fileOptions, "lba-localization.txt")
	saintFile := flags.SetString("saintFile", args, fileOptions, "")
	preyLimit := flags.SetInt("preyLimit", args, fileOptions, 100)

	options := parameters{
		database:  database,
		fdr:       fdr,
		namespace: namespace,
		minBaits:  minBaits,
		minFC:     minFC,
		outFile:   outFile,
		saintFile: saintFile,
		preyLimit: preyLimit,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.database == "" {
		messages = append(messages, "missing database file")
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
