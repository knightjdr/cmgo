package gradient

import (
	"bytes"
	"fmt"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeGradient(similarity []int, outfile string) {
	// Calculate height of svg. Their will be two pixels for each
	// gradient segment and a border of 1 pixel.
	height := (len(similarity) * 2) + 2

	colorMap := map[int]string{
		0: "#fff",
		1: "#ff8c66",
		2: "#cc3300",
	}

	var buffer bytes.Buffer

	// Create opening svg tag.
	buffer.WriteString(fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" height=\"%[1]d\" width=\"22\" viewBox=\"0 0 22 %[1]d\">\n", height))

	// Create gradient segments.
	for i, value := range similarity {
		y := (i * 2) + 1
		buffer.WriteString(fmt.Sprintf("\t<rect x=\"1\" y=\"%d\" height=\"2\" width=\"20\" fill=\"%s\" />\n", y, colorMap[value]))
	}

	// Create closing svg tag.
	buffer.WriteString(fmt.Sprintf("\t<rect x=\"0\" y=\"0\" height=\"%d\" width=\"22\" fill=\"none\" stroke=\"black\" stroke-width=\"1\" />\n", height))
	buffer.WriteString("</svg>\n")

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
