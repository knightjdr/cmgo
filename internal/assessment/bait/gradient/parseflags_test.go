package gradient

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
		"-baitList", "baits.txt",
		"-expectedLocalizations", "expected.txt",
		"-outFile", "out.txt",
	}
	fileOptions := map[string]interface{}{}
	wantArgs := parameters{
		baitList:              "baits.txt",
		expectedLocalizations: "expected.txt",
		outFile:               "out.txt",
	}
	args, err := parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required command line arguments are present")
	assert.Equal(t, wantArgs, args, "Should return arguments as options")

	// TEST2: return defaults when arguments missing.
	os.Args = []string{
		"cmd",
		"-baitList", "baits.txt",
		"-expectedLocalizations", "expected.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Equal(t, "bait-gradient.svg", args.outFile, "Should return default output file name")

	// TEST3: returns error when parameters are missing.
	os.Args = []string{
		"cmd",
	}
	wantErr := errors.New("missing ordered bait list; missing expected localizations")
	args, err = parseFlags(fileOptions)
	assert.NotNil(t, err, "Should return error when missing arguments")
	assert.Equal(t, wantErr, err, "Should return correct error message")

	// TEST4: reads parameters from file.
	os.Args = []string{
		"cmd",
	}
	fileOptions["baitList"] = "file-baits.txt"
	fileOptions["expectedLocalizations"] = "file-expected.txt"
	fileOptions["outFile"] = "file-out.txt"
	wantArgs = parameters{
		baitList:              "file-baits.txt",
		expectedLocalizations: "file-expected.txt",
		outFile:               "file-out.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required parameters are present")
	assert.Equal(t, wantArgs, args, "Should return file parameters as options")
}
