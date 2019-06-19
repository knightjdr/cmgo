package uv

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeSummary(data []map[string][]string, outfile string) {
	var buffer bytes.Buffer
	buffer.WriteString("rank\tknown\tunknown\tfraction\tgenes known\tgenes unknown\n")
	for i, assessment := range data {
		rank := i + 1
		known := len(assessment["known"])
		unknown := len(assessment["unknown"])
		total := known + unknown
		var fraction float64
		if total > 0 {
			fraction = float64(known) / float64(total)
		}
		sort.Strings(assessment["known"])
		sort.Strings(assessment["unknown"])
		buffer.WriteString(fmt.Sprintf("%d\t%d\t%d\t%0.2f\t%s\t%s\n", rank, known, unknown, fraction, strings.Join(assessment["known"], ", "), strings.Join(assessment["unknown"], ", ")))
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
