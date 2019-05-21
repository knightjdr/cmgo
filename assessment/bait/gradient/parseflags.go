package gradient

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/flags"
)

type parameters struct {
	baitList    string
	expectedLocalizations      string
	outFile          string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	baitList := flags.SetString("baitList", args, fileOptions, "")
	expectedLocalizations := flags.SetString("expectedLocalizations", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "bait-gradient.svg")

	// Copy arguments from options file.
	options := parameters{
		baitList:    baitList,
		expectedLocalizations:      expectedLocalizations,
		outFile:          outFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.baitList == "" {
		messages = append(messages, "missing ordered bait list")
	}
	if options.expectedLocalizations == "" {
		messages = append(messages, "missing expected localizations")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
