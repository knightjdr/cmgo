package nmfsafe

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
		"-goAnnotations", "annotations.gaf",
		"-goHierarchy", "hierarchy.obo",
		"-namespace", "BP",
		"-nmfLocalization", "nmf-localizations.txt",
		"-nmfSummary", "rank-summary.txt",
		"-outFile", "out.txt",
		"-outSummaryFile", "out-summary.txt",
		"-safeLocalization", "safe-localizations.txt",
		"-safeSummary", "domain-summary.txt",
	}
	fileOptions := map[string]interface{}{}
	wantArgs := parameters{
		goAnnotations:    "annotations.gaf",
		goHierarchy:      "hierarchy.obo",
		namespace:        "BP",
		nmfLocalization:  "nmf-localizations.txt",
		nmfSummary:       "rank-summary.txt",
		outFile:          "out.txt",
		outSummaryFile:   "out-summary.txt",
		safeLocalization: "safe-localizations.txt",
		safeSummary:      "domain-summary.txt",
	}
	args, err := parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required command line arguments are present")
	assert.Equal(t, wantArgs, args, "Should return arguments as options")

	// TEST2: return defaults when arguments missing.
	os.Args = []string{
		"cmd",
		"-goAnnotations", "annotations.gaf",
		"-goHierarchy", "hierarchy.obo",
		"-nmfLocalization", "nmf-localizations.txt",
		"-nmfSummary", "rank-summary.txt",
		"-safeLocalization", "safe-localizations.txt",
		"-safeSummary", "domain-summary.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Equal(t, "CC", args.namespace, "Should return default namespace")
	assert.Equal(t, "concordance.txt", args.outFile, "Should return default output file name")
	assert.Equal(t, "summary.txt", args.outSummaryFile, "Should return default summary output file name")

	// TEST3: returns error when parameters are missing.
	os.Args = []string{
		"cmd",
	}
	wantErr := errors.New("missing GO annotations (.gaf) file; missing GO hierarchy (.obo) file; missing NMF localization file; missing NMF rank summary file; missing SAFE localization file; missing SAFE rank summary file")
	args, err = parseFlags(fileOptions)
	assert.NotNil(t, err, "Should return error when missing arguments")
	assert.Equal(t, wantErr, err, "Should return correct error message")

	// TEST4: reads parameters from file.
	os.Args = []string{
		"cmd",
	}
	fileOptions["goAnnotations"] = "file-annotations.gaf"
	fileOptions["goHierarchy"] = "file-hierarchy.obo"
	fileOptions["namespace"] = "BP"
	fileOptions["nmfLocalization"] = "file-nmf-localizations.txt"
	fileOptions["nmfSummary"] = "file-rank-summary.txt"
	fileOptions["outFile"] = "file-out.txt"
	fileOptions["outSummaryFile"] = "file-out-summary.txt"
	fileOptions["safeLocalization"] = "file-safe-localizations.txt"
	fileOptions["safeSummary"] = "file-domain-summary.txt"
	wantArgs = parameters{
		goAnnotations:    "file-annotations.gaf",
		goHierarchy:      "file-hierarchy.obo",
		namespace:        "BP",
		nmfLocalization:  "file-nmf-localizations.txt",
		nmfSummary:       "file-rank-summary.txt",
		outFile:          "file-out.txt",
		outSummaryFile:   "file-out-summary.txt",
		safeLocalization: "file-safe-localizations.txt",
		safeSummary:      "file-domain-summary.txt",
	}
	args, err = parseFlags(fileOptions)
	assert.Nil(t, err, "Should not return an error when all required parameters are present")
	assert.Equal(t, wantArgs, args, "Should return file parameters as options")
}
