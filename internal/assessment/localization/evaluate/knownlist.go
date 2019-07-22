package evaluate

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func knownList(genes map[string]map[string]map[string]interface{}, outfile string) {
	// Determine write order for gene list.
	geneOrder := make([]string, 0)
	for gene := range genes {
		geneOrder = append(geneOrder, gene)
	}
	sort.Strings(geneOrder)

	var buffer bytes.Buffer
	buffer.WriteString("gene\tterm(s)\tID(s)\tknown\n")
	for _, gene := range geneOrder {
		terms := genes[gene]
		ids := make([]string, 0)
		idToName := make(map[string]string, 0)
		known := false
		for id, term := range terms {
			ids = append(ids, id)
			idToName[id] = term["name"].(string)
			if !known {
				known = term["known"].(bool)
			}
		}

		sort.Strings(ids)
		names := make([]string, len(ids))
		for i, id := range ids {
			names[i] = idToName[id]
		}
		buffer.WriteString(fmt.Sprintf("%s\t%s\t%s\t%t\n", gene, strings.Join(names, ";"), strings.Join(ids, ";"), known))
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
