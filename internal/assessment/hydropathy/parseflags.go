package hydropathy

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	bioplexFile string
	database    string
	fdr         float64
	saintFile   string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	bioplexFile := flags.SetString("bioplexFile", args, fileOptions, "")
	database := flags.SetString("database", args, fileOptions, "")
	fdr := flags.SetFloat("fdr", args, fileOptions, 0.01)
	saintFile := flags.SetString("saintFile", args, fileOptions, "")

	options := parameters{
		bioplexFile: bioplexFile,
		database:    database,
		fdr:         fdr,
		saintFile:   saintFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.bioplexFile == "" {
		messages = append(messages, "missing BioPlex File file")
	}
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
