package heatmap

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
		"-abundanceCap", "20",
		"-clusteringMethod", "average",
		"-compartmentSummary", "summary.txt",
		"-distanceMetric", "ward",
		"-enrichmentFile", "regions.txt",
		"-minAbundance", "5",
		"-outFile", "out.svg",
		"-pValue", "0.01",
	}
	fileOptions := map[string]interface{}{}
	wantArgs := parameters{
		abundanceCap:       20,
		clusteringMethod:   "average",
		compartmentSummary: "summary.txt",
		distanceMetric:     "ward",
		enrichmentFile:     "regions.txt",
		minAbundance:       5,
		outFile:            "out.svg",
		pValue:             0.01,
	}
	args, err := parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required command line arguments are present")
	assert.Equal(t, wantArgs, args, "Should return arguments as options")

	// TEST2: return defaults when arguments missing.
	os.Args = []string{
		"cmd",
		"-compartmentSummary", "summary.txt",
		"-enrichmentFile", "regions.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Equal(t, float64(10), args.abundanceCap, "Should return default abundance cap")
	assert.Equal(t, "complete", args.clusteringMethod, "Should return default clustering method")
	assert.Equal(t, "euclidean", args.distanceMetric, "Should return default distance metric")
	assert.Equal(t, float64(0), args.minAbundance, "Should return default minimum abundance")
	assert.Equal(t, "region-heatmap.svg", args.outFile, "Should return default outfile name")
	assert.Equal(t, 0.01, args.pValue, "Should return default pValue")

	// TEST3: returns error when parameters are missing.
	os.Args = []string{
		"cmd",
	}
	wantErr := errors.New("missing compartment summary file; missing enriched region file")
	args, err = parseFlags(fileOptions)
	assert.NotNil(t, err, "Should return error when missing arguments")
	assert.Equal(t, wantErr, err, "Should return correct error message")

	// TEST4: reads parameters from file.
	os.Args = []string{
		"cmd",
	}
	fileOptions["abundanceCap"] = 20
	fileOptions["clusteringMethod"] = "average"
	fileOptions["compartmentSummary"] = "file-summary.txt"
	fileOptions["distanceMetric"] = "ward"
	fileOptions["enrichmentFile"] = "file-regions.txt"
	fileOptions["minAbundance"] = 5
	fileOptions["outFile"] = "file-out.svg"
	fileOptions["pValue"] = 0.01
	wantArgs = parameters{
		abundanceCap:       20,
		clusteringMethod:   "average",
		compartmentSummary: "file-summary.txt",
		distanceMetric:     "ward",
		enrichmentFile:     "file-regions.txt",
		minAbundance:       5,
		outFile:            "file-out.svg",
		pValue:             0.01,
	}
	args, err = parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required parameters are present")
	assert.Equal(t, wantArgs, args, "Should return file parameters as options")
}
