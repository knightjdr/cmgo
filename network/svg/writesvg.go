package svg

import (
	"bytes"
	"fmt"
	"math"
	"sort"

	"github.com/knightjdr/cmgo/fs"
	"github.com/knightjdr/cmgo/read/nmf"
	"github.com/spf13/afero"
)

func writeSVG(coordinates map[string]coordinate, colors []string, localization nmf.NMFlocalization, plotWidth, plotHeight float64, outfile string) {
	// Determine order for outputting nodes.
	nodeOrder := make([]string, len(coordinates))
	i := 0
	for node := range coordinates {
		nodeOrder[i] = node
		i++
	}
	sort.Strings(nodeOrder)

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"%.2[1]f\" height=\"%.2[2]f\" viewbox=\"0 0 %.2[1]f %.2[2]f\">", plotWidth, plotHeight))
	buffer.WriteString("<g id=\"network__zoom\" transform=\"translate(0, 0) scale(1)\">")
	for _, node := range nodeOrder {
		position := coordinates[node]
		color := colors[localization[node].Rank]
		x := math.Round(position.X)
		y := math.Round(position.Y)
		buffer.WriteString(fmt.Sprintf("<circle cx=\"%.0f\" cy=\"%.0f\" r=\"4px\" fill=\"%s\" data-symbol=\"%s\" />", x, y, color, node))
	}
	buffer.WriteString("</g>")
	buffer.WriteString("</svg>")

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
