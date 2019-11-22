package knownbyrank

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeSummary(summary map[int]*rankSummary, outfile string) {
	var buffer bytes.Buffer

	writeHeader(&buffer)
	writeBody(&buffer, summary)

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}

func writeHeader(buffer *bytes.Buffer) {
	buffer.WriteString("prey rank\tproportion\tnumber of baits\tknown\tpairs\n")
}

func writeBody(buffer *bytes.Buffer, summary map[int]*rankSummary) {
	numberOfRanks := len(summary)

	for i := 1; i <= numberOfRanks; i++ {
		pairs := strings.Join(summary[i].Pairs, ", ")
		proportion := float64(summary[i].Known) / float64(summary[i].BaitNumber)
		buffer.WriteString(fmt.Sprintf("%d\t%0.2f\t%d\t%d\t%s\n", i, proportion, summary[i].BaitNumber, summary[i].Known, pairs))
	}
}
