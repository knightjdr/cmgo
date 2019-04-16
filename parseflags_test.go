package main

import (
	"errors"
	"os"
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var jsonText = `{
	"module": "test-module"
}`

func TestParseFlags(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/options.json",
		[]byte(jsonText),
		0444,
	)

	// Argument unmocking.
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// TEST1: reads arguments from JSON file.
	os.Args = []string{
		"cmd",
		"-options", "test/options.json",
	}
	wanted := map[string]interface{}{
		"module": "test-module",
	}
	options, err := parseFlags()
	assert.Nil(t, err, "Should not return an error with complete options file")
	assert.Equal(t, wanted, options, "Should return JSON as interface")

	// TEST2: overwrite options file parameter with command line.
	os.Args = []string{
		"cmd",
		"-module", "overwrite-test",
		"-options", "test/options.json",
	}
	wanted = map[string]interface{}{
		"module": "overwrite-test",
	}
	options, _ = parseFlags()
	assert.Equal(t, wanted, options, "Should overwrite options file parameters")

	// TEST3: no module specified
	os.Args = []string{
		"cmd",
	}
	wantedErr := errors.New("no module specified")
	_, err = parseFlags()
	assert.NotNil(t, err, "Should return an error when analysis module not specified")
	assert.Equal(t, wantedErr, err, "Should return error message")
}
