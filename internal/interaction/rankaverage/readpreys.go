package rankaverage

import "github.com/knightjdr/cmgo/internal/pkg/read/list"

func readPreys(filename string) []string {
	fileData := list.CSV(filename, '\t')

	preys := make([]string, len(fileData))
	for i, line := range fileData {
		preys[i] = line["gene"]
	}

	return preys
}
