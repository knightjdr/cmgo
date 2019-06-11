package crapome

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
		"-baitFiles", "bait-task1.dat;bait-task2.dat",
		"-crapomeID", "cc.txt",
		"-interactionFiles", "inter-task1.dat;inter-task2.dat",
		"-outFile", "out.txt",
		"-preyFiles", "prey-task1.dat;prey-task2.dat",
	}
	fileOptions := map[string]interface{}{}
	wantArgs := parameters{
		baitFiles:        []string{"bait-task1.dat", "bait-task2.dat"},
		crapomeIDfile:    "cc.txt",
		interactionFiles: []string{"inter-task1.dat", "inter-task2.dat"},
		outFile:          "out.txt",
		preyFiles:        []string{"prey-task1.dat", "prey-task2.dat"},
	}
	args, err := parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required command line arguments are present")
	assert.Equal(t, wantArgs, args, "Should return arguments as options")

	// TEST2: return defaults when arguments missing.
	os.Args = []string{
		"cmd",
		"-baitFiles", "bait-task1.dat;bait-task2.dat",
		"-crapomeID", "cc.txt",
		"-interactionFiles", "inter-task1.dat;inter-task2.dat",
		"-preyFiles", "prey-task1.dat;prey-task2.dat",
	}
	args, err = parseFlags(fileOptions)
	assert.Equal(t, "crapome-matrix.txt", args.outFile, "Should return default outfile name")

	// TEST3: returns error when parameters are missing.
	os.Args = []string{
		"cmd",
	}
	wantErr := errors.New("missing bait.dat file; missing inter.dat file; missing prey.dat file")
	args, err = parseFlags(fileOptions)
	assert.NotNil(t, err, "Should return error when missing arguments")
	assert.Equal(t, wantErr, err, "Should return correct error message")

	// TEST4: reads parameters from file.
	os.Args = []string{
		"cmd",
	}
	fileOptions["baitFiles"] = "file-bait-task1.dat;file-bait-task2.dat"
	fileOptions["crapomeID"] = "file-cc.txt"
	fileOptions["interactionFiles"] = "file-inter-task1.dat;file-inter-task2.dat"
	fileOptions["outFile"] = "file-out.txt"
	fileOptions["preyFiles"] = "file-prey-task1.dat;file-prey-task2.dat"
	wantArgs = parameters{
		baitFiles:        []string{"file-bait-task1.dat", "file-bait-task2.dat"},
		crapomeIDfile:    "file-cc.txt",
		interactionFiles: []string{"file-inter-task1.dat", "file-inter-task2.dat"},
		outFile:          "file-out.txt",
		preyFiles:        []string{"file-prey-task1.dat", "file-prey-task2.dat"},
	}
	args, err = parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required parameters are present")
	assert.Equal(t, wantArgs, args, "Should return file parameters as options")
}
