package tsnecytoscape

import (
	"math"

	"github.com/knightjdr/cmgo/internal/pkg/read/tsne"
)

func networkTransformation(coordinates map[string]tsne.Coordinate, scaleWidth float64) map[string]float64 {
	transformation := map[string]float64{
		"scale":      1,
		"translateX": math.MaxFloat64,
		"translateY": math.MaxFloat64,
	}
	maxX := -math.MaxFloat64

	for _, position := range coordinates {
		if position.X > maxX {
			maxX = position.X
		}
		if position.X < transformation["translateX"] {
			transformation["translateX"] = position.X
		}
		if position.Y < transformation["translateY"] {
			transformation["translateY"] = position.Y
		}
	}

	transformation["scale"] = scaleWidth / (maxX - transformation["translateX"])
	transformation["translateX"] = -1 * transformation["translateX"]
	transformation["translateY"] = -1 * transformation["translateY"]
	return transformation
}
