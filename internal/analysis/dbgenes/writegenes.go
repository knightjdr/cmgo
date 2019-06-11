package dbgenes

import (
	"bytes"
	"fmt"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/slice"
	"github.com/spf13/afero"
)

func writeGenes(genes []string, outfile string) {
	sorted := slice.SortStringsCaseInsensitive(genes)

	var buffer bytes.Buffer
	for _, gene := range sorted {
		buffer.WriteString(fmt.Sprintf("%s\n", gene))
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
