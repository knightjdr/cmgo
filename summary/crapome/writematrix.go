package crapome

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/knightjdr/cmgo/fs"
	"github.com/knightjdr/cmgo/read"
	"github.com/knightjdr/cmgo/stats"
	"github.com/spf13/afero"
)

func writeMatrix(data map[string]map[string]int, baits []read.BaitDatRow, preyMap map[string]*preyDefinition, outfile string) {
	// Order preys alphabetically by name. First get gene name for each prey
	// from map. Sort by gene name, then map that back to accession.
	preyOrder := make([]string, len(preyMap))
	reverseMap := make(map[string]string, len(preyMap))
	i := 0
	for prey, preyDefinition := range preyMap {
		preyOrder[i] = preyDefinition.Name
		reverseMap[preyDefinition.Name] = prey
		i++
	}
	sort.Strings(preyOrder)
	for i, name := range preyOrder {
		preyOrder[i] = reverseMap[name]
	}

	// Order samples alphabetically.
	sampleOrder := make([]string, len(baits))
	i = 0
	for _, bait := range baits {
		sampleOrder[i] = bait.ID
		i++
	}
	sort.Strings(sampleOrder)

	// Create file header with sample IDs.
	var buffer bytes.Buffer
	buffer.WriteString("GENE\tREFSEQ_ID\tENTREZ_ID\tAVE_SC\tNUM_EXPT")
	for _, id := range sampleOrder {
		buffer.WriteString(fmt.Sprintf("\t%s", id))
	}
	buffer.WriteString("\n")

	// Write prey rows.
	for _, accession := range preyOrder {
		numExpt := len(data[accession])
		spectralCounts := make([]int, len(data[accession]))
		i = 0
		for _, spec := range data[accession] {
			spectralCounts[i] = spec
			i++
		}
		aveSC := stats.MeanInt(spectralCounts)

		buffer.WriteString(fmt.Sprintf("%s\t%s\t%d\t%0.2f\t%d", preyMap[accession].Name, accession, preyMap[accession].GeneID, aveSC, numExpt))
		for _, id := range sampleOrder {
			var spec int
			if _, ok := data[accession][id]; ok {
				spec = data[accession][id]
			}
			buffer.WriteString(fmt.Sprintf("\t%d", spec))
		}
		buffer.WriteString("\n")
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
