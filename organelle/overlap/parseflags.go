package overlap

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/flags"
)

func parseFlags(fileOptions map[string]interface{}) (map[string]string, error) {
	args := flags.Parse()
	compartmentFile := flags.ConvertString(args["compartmentFile"])
	outFile := flags.ConvertString(args["outFile"])
	similarityFile := flags.ConvertString(args["similarityFile"])

	// Copy arguments from options file.
	options := map[string]string{}
	if fileOptions["compartmentFile"] != nil {
		options["compartmentFile"] = fileOptions["compartmentFile"].(string)
	}
	if fileOptions["outFile"] != nil {
		options["outFile"] = fileOptions["outFile"].(string)
	}
	if fileOptions["similarityFile"] != nil {
		options["similarityFile"] = fileOptions["similarityFile"].(string)
	}

	// Overwrite options file arguments if specified
	if compartmentFile != "" {
		options["compartmentFile"] = compartmentFile
	}
	if outFile != "" {
		options["outFile"] = outFile
	}
	if similarityFile != "" {
		options["similarityFile"] = similarityFile
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options["compartmentFile"] == "" {
		messages = append(messages, "missing JSON file with list of compartments to compare")
	}
	if options["similarityFile"] == "" {
		messages = append(messages, "missing search result peptide file")
	}
	if options["outFile"] == "" {
		options["outFile"] = "organelle-comparison.txt"
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
