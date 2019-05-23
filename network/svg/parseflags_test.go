package svg

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFlags(t *testing.T) {
	// Argument unmocking.
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// TEST1: return options from command line arguments.
	os.Args = []string{
		"cmd",
		"-colorList", "colors.txt",
		"-localizations", "localizations.txt",
		"-nodeCoordinates", "coordinates.txt",
		"-outFile", "out.svg",
	}
	fileOptions := map[string]interface{}{}
	wantArgs := parameters{
		colorList:       "colors.txt",
		localizations:   "localizations.txt",
		nodeCoordinates: "coordinates.txt",
		outFile:         "out.svg",
	}
	args, err := parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required command line arguments are present")
	assert.Equal(t, wantArgs, args, "Should return arguments as options")

	// TEST2: return defaults when arguments missing.
	os.Args = []string{
		"cmd",
		"-colorList", "colors.txt",
		"-localizations", "localizations.txt",
		"-nodeCoordinates", "coordinates.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Equal(t, "network.svg", args.outFile, "Should return default outfile name")

	// TEST3: returns error when parameters are missing.
	os.Args = []string{
		"cmd",
	}
	wantErr := errors.New("missing color list; missing localization file; missing node coordinates")
	args, err = parseFlags(fileOptions)
	assert.NotNil(t, err, "Should return error when missing arguments")
	assert.Equal(t, wantErr, err, "Should return correct error message")

	// TEST4: reads parameters from file.
	os.Args = []string{
		"cmd",
	}
	fileOptions["colorList"] = "file-colors.txt"
	fileOptions["localizations"] = "file-localizations.txt"
	fileOptions["nodeCoordinates"] = "file-coordinates.txt"
	fileOptions["outFile"] = "file-out.svg"
	wantArgs = parameters{
		colorList:       "file-colors.txt",
		localizations:   "file-localizations.txt",
		nodeCoordinates: "file-coordinates.txt",
		outFile:         "file-out.svg",
	}
	args, err = parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required parameters are present")
	assert.Equal(t, wantArgs, args, "Should return file parameters as options")
}
