package rankaverage

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeAverages(rankAverages map[string]preySummary, outfile string) {
	var buffer bytes.Buffer

	writeHeader(&buffer)
	writeBody(&buffer, rankAverages)

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}

func writeHeader(buffer *bytes.Buffer) {
	buffer.WriteString("prey\tmean\tsd\tinteraction ranks\n")
}

func writeBody(buffer *bytes.Buffer, rankAverages map[string]preySummary) {
	outputOrder := getOutputOder(rankAverages)

	for _, prey := range outputOrder {
		summary := rankAverages[prey]
		rankString := joinRanks(summary.ranks)
		buffer.WriteString(fmt.Sprintf("%s\t%0.3f\t%0.3f\t\"%s\"\n", prey, summary.mean, summary.sd, rankString))
	}
}

func getOutputOder(rankAverages map[string]preySummary) []string {
	keys := make([]string, len(rankAverages))

	i := 0
	for key := range rankAverages {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	return keys
}

func joinRanks(ranks []int) string {
	strs := make([]string, len(ranks))
	for i, value := range ranks {
		strs[i] = strconv.Itoa(value)
	}
	return strings.Join(strs, ",")
}
