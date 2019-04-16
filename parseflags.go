package main

import (
	"encoding/json"
	"errors"

	"github.com/knightjdr/cmgo/flags"
	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
)

func parseFlags() (map[string]interface{}, error) {
	args := flags.Parse()
	module := flags.ConvertString(args["module"])
	optionsFile := flags.ConvertString(args["options"])

	var err error
	var options map[string]interface{}

	// Read options from file if specified.
	if optionsFile != "" {
		jsonFile, _ := afero.ReadFile(fs.Instance, optionsFile)
		var jsonData interface{}
		err = json.Unmarshal(jsonFile, &jsonData)
		options, _ = jsonData.(map[string]interface{})
	}

	// Overwite file parameter for module with command line argument if specified.
	if module != "" {
		options["module"] = module
	}

	// Check if an analysis module has been specified.
	if options["module"] == nil {
		err = errors.New("no module specified")
	}

	return options, err
}
