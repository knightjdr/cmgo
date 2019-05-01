package dbgenes

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/flags"
)

type parameters struct {
	database        string
	outFile          string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	database := flags.SetString("database", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "db-genes.txt")

	// Copy arguments from options file.
	options := parameters{
		database:        database,
		outFile:          outFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.database == "" {
		messages = append(messages, "missing FASTA database file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
