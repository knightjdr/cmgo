package overlap

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
		"-compartmentFile", "compartments.json",
		"-outFile", "out.txt",
		"-similarityFile", "sim.txt",
	}
	fileOptions := map[string]interface{}{}
	wantArgs := map[string]string{
		"compartmentFile": "compartments.json",
		"outFile":         "out.txt",
		"similarityFile":  "sim.txt",
	}
	args, err := parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required command line arguments are present")
	assert.Equal(t, wantArgs, args, "Should return arguments as options")

	// TEST2: return default name for out file.
	os.Args = []string{
		"cmd",
		"-compartmentFile", "compartments.json",
		"-similarityFile", "sim.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Equal(t, "organelle-comparison.txt", args["outFile"], "Should return default outfile name")

	// TEST3: returns error when parameters are missing.
	os.Args = []string{
		"cmd",
	}
	wantErr := errors.New("missing JSON file with list of compartments to compare; missing search result peptide file")
	args, err = parseFlags(fileOptions)
	assert.NotNil(t, err, "Should return error when missing arguments")
	assert.Equal(t, wantErr, err, "Should return correct error message")

	// TEST4: reads parameters from file.
	os.Args = []string{
		"cmd",
	}
	fileOptions["compartmentFile"] = "file-compartments.json"
	fileOptions["outFile"] = "file-out.txt"
	fileOptions["similarityFile"] = "file-sim.txt"
	wantArgs = map[string]string{
		"compartmentFile": "file-compartments.json",
		"outFile":         "file-out.txt",
		"similarityFile":  "file-sim.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required parameters are present")
	assert.Equal(t, wantArgs, args, "Should return file parameters as options")
}
