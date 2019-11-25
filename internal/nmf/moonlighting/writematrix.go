package moonlighting

import (
	"bytes"
	"fmt"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeRankMoonlightingMatrix(matrix [][]int, outfile string) {
	var buffer bytes.Buffer

	writeMatrixHeader(&buffer, len(matrix))
	writeMatrixBody(&buffer, matrix)

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}

func writeMatrixHeader(buffer *bytes.Buffer, dimension int) {
	for i := 0; i < dimension; i++ {
		buffer.WriteString(fmt.Sprintf("\t%d", i+1))
	}
	buffer.WriteString("\n")
}

func writeMatrixBody(buffer *bytes.Buffer, matrix [][]int) {
	for rowIndex, row := range matrix {
		buffer.WriteString(fmt.Sprintf("%d", rowIndex+1))
		for _, value := range row {
			buffer.WriteString(fmt.Sprintf("\t%d", value))
		}
		buffer.WriteString("\n")
	}
}
