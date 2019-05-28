package subset

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
		"-basisMatrix", "basis.csv",
		"-clusteringMethod", "average",
		"-distanceMetric", "ward",
		"-minAbundance", "5",
		"-minNMFScore", "0.15",
		"-outFile", "out.svg",
		"-ranks1", "1,5",
		"-ranks2", "7",
		"-specificity", "3",
		"-threshold", "0.2",
	}
	fileOptions := map[string]interface{}{}
	wantArgs := parameters{
		abundanceCap:     20,
		basisMatrix:      "basis.csv",
		clusteringMethod: "average",
		distanceMetric:   "ward",
		minAbundance:     5,
		minNMFScore:      0.15,
		outFile:          "out.svg",
		ranks1:           []string{"1", "5"},
		ranks2:           []string{"7"},
		specificity:      3,
		threshold:        0.2,
	}
	args, err := parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required command line arguments are present")
	assert.Equal(t, wantArgs, args, "Should return arguments as options")

	// TEST2: return defaults when arguments missing.
	os.Args = []string{
		"cmd",
		"-basisMatrix", "basis.csv",
		"-ranks1", "1,5",
		"-ranks2", "7",
	}
	args, err = parseFlags(fileOptions)
	assert.Equal(t, float64(10), args.abundanceCap, "Should return default abundance cap")
	assert.Equal(t, "complete", args.clusteringMethod, "Should return default clustering method")
	assert.Equal(t, "euclidean", args.distanceMetric, "Should return default distance metric")
	assert.Equal(t, float64(0), args.minAbundance, "Should return default minimum abundance")
	assert.Equal(t, float64(0), args.minNMFScore, "Should return default minimum NMF score")
	assert.Equal(t, "basis-subset.svg", args.outFile, "Should return default outfile name")
	assert.Equal(t, float64(2), args.specificity, "Should return default specificity score")
	assert.Equal(t, 0.5, args.threshold, "Should return default threshold")

	// TEST3: returns error when parameters are missing.
	os.Args = []string{
		"cmd",
	}
	wantErr := errors.New("missing basis matrix file; missing first compartment to check; missing second compartment to check")
	args, err = parseFlags(fileOptions)
	assert.NotNil(t, err, "Should return error when missing arguments")
	assert.Equal(t, wantErr, err, "Should return correct error message")

	// TEST4: reads parameters from file.
	os.Args = []string{
		"cmd",
	}
	fileOptions["abundanceCap"] = 20
	fileOptions["basisMatrix"] = "file-basis.csv"
	fileOptions["clusteringMethod"] = "average"
	fileOptions["distanceMetric"] = "ward"
	fileOptions["minAbundance"] = 5
	fileOptions["minNMFScore"] = 0.15
	fileOptions["outFile"] = "file-out.svg"
	fileOptions["ranks1"] = "1,5"
	fileOptions["ranks2"] = "7"
	fileOptions["specificity"] = 3
	fileOptions["threshold"] = 0.2
	wantArgs = parameters{
		abundanceCap:     20,
		basisMatrix:      "file-basis.csv",
		clusteringMethod: "average",
		distanceMetric:   "ward",
		minAbundance:     5,
		minNMFScore:      0.15,
		outFile:          "file-out.svg",
		ranks1:           []string{"1", "5"},
		ranks2:           []string{"7"},
		specificity:      3,
		threshold:        0.2,
	}
	args, err = parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required parameters are present")
	assert.Equal(t, wantArgs, args, "Should return file parameters as options")
}
