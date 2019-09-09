package dbgenes

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	ncbigene string
	outFile  string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	ncbigene := flags.SetString("ncbigene", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "db-genes.txt")

	// Copy arguments from options file.
	options := parameters{
		ncbigene: ncbigene,
		outFile:  outFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.ncbigene == "" {
		messages = append(messages, "missing FASTA ncbigene file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
