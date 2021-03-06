package crapome

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/stats"
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	"github.com/spf13/afero"
)

func writeMatrix(data map[string]map[string]int, baits []saint.BaitDatRow, preyMap map[string]string, preyOrder []string, idToCCmap map[int]string, outfile string) {
	// Order samples alphabetically by CC number. Create a map for
	// for retrieving bait ID from CC number since data is stored
	// by bait ID.
	ccToIDmap := make(map[string]string, 0)
	sampleOrder := make([]string, len(baits))
	i := 0
	for _, bait := range baits {
		id, _ := strconv.Atoi(strings.Split(bait.ID, "_")[1])
		ccToIDmap[idToCCmap[id]] = bait.ID
		sampleOrder[i] = idToCCmap[id]
		i++
	}
	sort.Strings(sampleOrder)

	// Create file header with sample IDs.
	var buffer bytes.Buffer
	buffer.WriteString("GENE\tREFSEQ_ID\tAVE_SC\tNUM_EXPT")
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

		buffer.WriteString(fmt.Sprintf("%s\t%s\t%0.2f\t%d", preyMap[accession], accession, aveSC, numExpt))
		for _, cc := range sampleOrder {
			id := ccToIDmap[cc]
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
