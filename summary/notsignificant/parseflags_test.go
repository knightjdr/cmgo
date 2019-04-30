package notsignificant

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
		"-fdr", "0.05",
		"-outFile", "out.txt",
		"-saint", "saint.txt",
	}
	fileOptions := map[string]interface{}{}
	wantArgs := parameters{
		fdr:       0.05,
		outFile:   "out.txt",
		saintFile: "saint.txt",
	}
	args, err := parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required command line arguments are present")
	assert.Equal(t, wantArgs, args, "Should return arguments as options")

	// TEST2: return defaults when arguments missing.
	os.Args = []string{
		"cmd",
		"-saint", "saint.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Equal(t, 0.01, args.fdr, "Should return default FDR")
	assert.Equal(t, "not-significant.txt", args.outFile, "Should return default outfile name")

	// TEST3: returns error when parameters are missing.
	os.Args = []string{
		"cmd",
	}
	wantErr := errors.New("missing SAINT file")
	args, err = parseFlags(fileOptions)
	assert.NotNil(t, err, "Should return error when missing arguments")
	assert.Equal(t, wantErr, err, "Should return correct error message")

	// TEST4: reads parameters from file.
	os.Args = []string{
		"cmd",
	}
	fileOptions["fdr"] = 0.05
	fileOptions["outFile"] = "file-out.txt"
	fileOptions["saint"] = "file-saint.txt"
	wantArgs = parameters{
		fdr:       0.05,
		outFile:   "file-out.txt",
		saintFile: "file-saint.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required parameters are present")
	assert.Equal(t, wantArgs, args, "Should return file parameters as options")
}
