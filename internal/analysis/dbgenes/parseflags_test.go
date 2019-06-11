package dbgenes

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
		"-database", "sequence.fasta",
		"-outFile", "out.txt",
	}
	fileOptions := map[string]interface{}{}
	wantArgs := parameters{
		database: "sequence.fasta",
		outFile:  "out.txt",
	}
	args, err := parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required command line arguments are present")
	assert.Equal(t, wantArgs, args, "Should return arguments as options")

	// TEST2: return defaults when arguments missing.
	os.Args = []string{
		"cmd",
		"-database", "sequence.fasta",
	}
	args, err = parseFlags(fileOptions)
	assert.Equal(t, "db-genes.txt", args.outFile, "Should return default outfile name")

	// TEST3: returns error when parameters are missing.
	os.Args = []string{
		"cmd",
	}
	wantErr := errors.New("missing FASTA database file")
	args, err = parseFlags(fileOptions)
	assert.NotNil(t, err, "Should return error when missing arguments")
	assert.Equal(t, wantErr, err, "Should return correct error message")

	// TEST4: reads parameters from file.
	os.Args = []string{
		"cmd",
	}
	fileOptions["database"] = "file-sequences.fasta"
	fileOptions["outFile"] = "file-out.txt"
	wantArgs = parameters{
		database: "file-sequences.fasta",
		outFile:  "file-out.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required parameters are present")
	assert.Equal(t, wantArgs, args, "Should return file parameters as options")
}
