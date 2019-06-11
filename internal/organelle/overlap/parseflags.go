package overlap

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

func parseFlags(fileOptions map[string]interface{}) (map[string]string, error) {
	args := flags.Parse()
	compartmentFile := flags.SetString("compartmentFile", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "organelle-comparison.txt")
	similarityFile := flags.SetString("similarityFile", args, fileOptions, "")

	options := map[string]string{
		"compartmentFile": compartmentFile,
		"outFile":         outFile,
		"similarityFile":  similarityFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options["compartmentFile"] == "" {
		messages = append(messages, "missing JSON file with list of compartments to compare")
	}
	if options["similarityFile"] == "" {
		messages = append(messages, "missing search result peptide file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
