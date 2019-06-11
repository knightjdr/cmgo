package crapome

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	baitFiles        []string
	crapomeIDfile    string
	interactionFiles []string
	outFile          string
	preyFiles        []string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	baitFile := flags.SetString("baitFiles", args, fileOptions, "")
	crapomeIDfile := flags.SetString("crapomeID", args, fileOptions, "")
	interactionFile := flags.SetString("interactionFiles", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "crapome-matrix.txt")
	preyFile := flags.SetString("preyFiles", args, fileOptions, "")

	// Copy arguments from options file.
	options := parameters{
		baitFiles:        strings.Split(baitFile, ";"),
		crapomeIDfile:    crapomeIDfile,
		interactionFiles: strings.Split(interactionFile, ";"),
		outFile:          outFile,
		preyFiles:        strings.Split(preyFile, ";"),
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if len(options.baitFiles) == 0 || options.baitFiles[0] == "" {
		messages = append(messages, "missing bait.dat file")
	}
	if len(options.interactionFiles) == 0 || options.interactionFiles[0] == "" {
		messages = append(messages, "missing inter.dat file")
	}
	if len(options.preyFiles) == 0 || options.preyFiles[0] == "" {
		messages = append(messages, "missing prey.dat file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
