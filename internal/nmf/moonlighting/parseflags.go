package moonlighting

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	dissimilarityFile string
	minimumNmfScore   float64
	nmfBasis          string
	outFile           string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	dissimilarityFile := flags.SetString("dissimilarityFile", args, fileOptions, "")
	minimumNmfScore := flags.SetFloat("minimumNmfScore", args, fileOptions, 0.15)
	nmfBasis := flags.SetString("nmfBasis", args, fileOptions, "")
	outFile := flags.SetString("outFile", args, fileOptions, "moonlighting.txt")

	// Copy arguments from options file.
	options := parameters{
		dissimilarityFile: dissimilarityFile,
		minimumNmfScore:   minimumNmfScore,
		nmfBasis:          nmfBasis,
		outFile:           outFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.dissimilarityFile == "" {
		messages = append(messages, "missing dissimilarity file")
	}
	if options.nmfBasis == "" {
		messages = append(messages, "missing NMF basis file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
