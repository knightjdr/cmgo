package evaluate

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

type localizationSummary struct {
	known int
	name  string
	total int
}

func summarize(genes map[string]map[string]map[string]interface{}, outfile string) {
	ids := make([]string, 0)
	termData := make(map[string]*localizationSummary, 0)
	knownCount := 0
	for _, terms := range genes {
		known := false
		for id, term := range terms {
			goID := id
			if goID == "" {
				goID = "-"
			}
			if !known && term["known"].(bool) {
				known = true
				knownCount++
			}
			if _, ok := termData[goID]; !ok {
				ids = append(ids, goID)
				var name string
				if term["name"].(string) != "" {
					name = term["name"].(string)
				} else {
					name = "unknown"
				}
				termData[goID] = &localizationSummary{
					name: name,
				}
			}
			termData[goID].total++
			if term["known"].(bool) {
				termData[goID].known++
			}
		}
	}

	totalGenes := len(genes)
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("total genes: %d\n", totalGenes))
	buffer.WriteString(fmt.Sprintf("known genes: %d\n", knownCount))
	buffer.WriteString(fmt.Sprintf("fraction known: %0.4f\n", float64(knownCount)/float64(totalGenes)))

	buffer.WriteString("\nGO ID\tterm\ttotal\tknown\tfraction\n")
	sort.Strings(ids)
	for _, id := range ids {
		data := termData[id]
		fraction := float64(data.known) / float64(data.total)
		buffer.WriteString(fmt.Sprintf("%s\t%s\t%d\t%d\t%0.4f\n", id, data.name, data.total, data.known, fraction))
	}

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
