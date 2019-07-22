package localize

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	enrichment     string
	localization   string
	outFilePrimary string
	outFileProfile string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	enrichment := flags.SetString("enrichment", args, fileOptions, "")
	localization := flags.SetString("localization", args, fileOptions, "")
	outFilePrimary := flags.SetString("outFilePrimary", args, fileOptions, "lba-primary.txt")
	outFileProfile := flags.SetString("outFileProfile", args, fileOptions, "lba-profile.txt")

	options := parameters{
		enrichment:     enrichment,
		localization:   localization,
		outFilePrimary: outFilePrimary,
		outFileProfile: outFileProfile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.enrichment == "" {
		messages = append(messages, "missing enrichment file")
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
