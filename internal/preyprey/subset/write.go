package subset

import (
	"bytes"
	"fmt"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func write(lines [][]string, params, outfile string) {
	var buffer bytes.Buffer
	buffer.WriteString("row\tcolumn\tvalue\tparams\n")

	for i, line := range lines {
		if i == 0 {
			buffer.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", line[0], line[1], line[2], params))
		} else {
			buffer.WriteString(fmt.Sprintf("%s\t%s\t%s\n", line[0], line[1], line[2]))
		}
	}

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
