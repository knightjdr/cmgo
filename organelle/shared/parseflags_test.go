package shared

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
		"-fdr", "0.01",
		"-outFile", "out.txt",
		"-regionFile", "regions.txt",
		"-saintFile", "saint.txt",
	}
	fileOptions := map[string]interface{}{}
	wantArgs := parameters{
		compartmentFile: "compartments.json",
		fdr:             0.01,
		outFile:         "out.txt",
		regionFile:      "regions.txt",
		saintFile:       "saint.txt",
	}
	args, err := parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required command line arguments are present")
	assert.Equal(t, wantArgs, args, "Should return arguments as options")

	// TEST2: return default name for out file and FDR.
	os.Args = []string{
		"cmd",
		"-compartmentFile", "compartments.json",
		"-regionFile", "regions.txt",
		"-saintFile", "saint.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Equal(t, float64(0), args.fdr, "Should return default fdr")
	assert.Equal(t, "organelle-shared.txt", args.outFile, "Should return default outfile name")

	// TEST3: returns error when parameters are missing.
	os.Args = []string{
		"cmd",
	}
	wantErr := errors.New("missing JSON file with list of compartments to compare; missing region file; missing SAINT file")
	args, err = parseFlags(fileOptions)
	assert.NotNil(t, err, "Should return error when missing arguments")
	assert.Equal(t, wantErr, err, "Should return correct error message")

	// TEST4: reads parameters from file.
	os.Args = []string{
		"cmd",
	}
	fileOptions["compartmentFile"] = "file-compartments.json"
	fileOptions["fdr"] = 0.01
	fileOptions["outFile"] = "file-out.txt"
	fileOptions["regionFile"] = "file-regions.txt"
	fileOptions["saintFile"] = "file-saint.txt"
	wantArgs = parameters{
		compartmentFile: "file-compartments.json",
		fdr:             0.01,
		outFile:         "file-out.txt",
		regionFile:      "file-regions.txt",
		saintFile:       "file-saint.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required parameters are present")
	assert.Equal(t, wantArgs, args, "Should return file parameters as options")
}
