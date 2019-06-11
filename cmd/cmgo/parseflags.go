package main

import (
	"encoding/json"
	"errors"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/flags"
	"github.com/spf13/afero"
)

func parseFlags() (map[string]interface{}, error) {
	args := flags.Parse()
	module := flags.SetString("module", args, map[string]interface{}{}, "")
	optionsFile := flags.SetString("options", args, map[string]interface{}{}, "")

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
