package rankmetrics

import (
	"errors"
	"strings"

	"github.com/knightjdr/cmgo/pkg/flags"
)

type parameters struct {
	fasta        string
	fdr          float64
	outFile      string
	saint        string
	turnoverFile string
}

func parseFlags(fileOptions map[string]interface{}) (parameters, error) {
	args := flags.Parse()
	fasta := flags.SetString("fasta", args, fileOptions, "")
	fdr := flags.SetFloat("fdr", args, fileOptions, 0.01)
	outFile := flags.SetString("outFile", args, fileOptions, "turnover-by-rank.txt")
	saint := flags.SetString("saint", args, fileOptions, "")
	turnoverFile := flags.SetString("turnoverFile", args, fileOptions, "")

	// Copy arguments from options file.
	options := parameters{
		fasta:        fasta,
		fdr:          fdr,
		outFile:      outFile,
		saint:        saint,
		turnoverFile: turnoverFile,
	}

	// Check for missing arguments.
	messages := make([]string, 0)
	if options.fasta == "" {
		messages = append(messages, "missing fasta database")
	}
	if options.saint == "" {
		messages = append(messages, "missing SAINT file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	var err error
	if errorString != "" {
		err = errors.New(errorString)
	}

	return options, err
}
