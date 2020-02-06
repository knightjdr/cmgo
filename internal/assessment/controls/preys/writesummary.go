package preys

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/slice"
	"github.com/spf13/afero"
)

func writeSummary(summary map[string]*preyInteraction, outFile string) {
	var buffer bytes.Buffer

	writeHeader(&buffer)
	writeBody(&buffer, summary)

	afero.WriteFile(fs.Instance, outFile, buffer.Bytes(), 0644)
}

func writeHeader(buffer *bytes.Buffer) {
	header := "prey\taverage\tmax\tmin\t" +
		"average BirA-FLAG\tmax BirA-FLAG\tmin BirA-FLAG\t" +
		"average BirA-GFP\tmax BirA-GFP\tmin BirA-GFP\t" +
		"average empty\tmax empty\tmin empty\t" +
		"BirA-Flag\tBirA-FGFP\tempty\n"
	buffer.WriteString(header)
}

func writeBody(buffer *bytes.Buffer, summary map[string]*preyInteraction) {
	outputOrder := orderPreys(summary)

	for _, prey := range outputOrder {
		buffer.WriteString(
			fmt.Sprintf(
				"%s\t%0.2f\t%d\t%d\t%0.2f\t%d\t%d\t%0.2f\t%d\t%d\t%0.2f\t%d\t%d\t%s\t%s\t%s\n",
				prey,
				summary[prey].Average.Overall,
				summary[prey].Max.Overall,
				summary[prey].Min.Overall,
				summary[prey].Average.BirAFlag,
				summary[prey].Max.BirAFlag,
				summary[prey].Min.BirAFlag,
				summary[prey].Average.BirAGFP,
				summary[prey].Max.BirAGFP,
				summary[prey].Min.BirAGFP,
				summary[prey].Average.Empty,
				summary[prey].Max.Empty,
				summary[prey].Min.Empty,
				slice.JoinInts(summary[prey].BirAFlag, "|"),
				slice.JoinInts(summary[prey].BirAGFP, "|"),
				slice.JoinInts(summary[prey].Empty, "|"),
			),
		)
	}
}

type orderFields struct {
	Average float64
	Prey    string
}

func orderPreys(summary map[string]*preyInteraction) []string {
	noPreys := len(summary)
	preyFields := make([]orderFields, noPreys)

	i := 0
	for prey, preySummary := range summary {
		preyFields[i] = orderFields{
			Average: preySummary.Average.Overall,
			Prey:    prey,
		}
		i++
	}

	sort.Slice(preyFields, func(i, j int) bool {
		return preyFields[i].Average > preyFields[j].Average
	})

	preys := make([]string, noPreys)
	for i, fields := range preyFields {
		preys[i] = fields.Prey
	}

	return preys
}
