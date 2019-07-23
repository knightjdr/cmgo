package svg

import (
	"math"

	"github.com/knightjdr/cmgo/internal/pkg/read/tsne"
)

func scaleCoordinates(coordinates map[string]tsne.Coordinate, scaleWidth float64, radius, borderWidth int) (map[string]tsne.Coordinate, float64) {
	maxX := -math.MaxFloat64
	maxY := -math.MaxFloat64
	minX := math.MaxFloat64
	minY := math.MaxFloat64

	for _, position := range coordinates {
		if position.X < minX {
			minX = position.X
		} else if position.X > maxX {
			maxX = position.X
		}

		if position.Y < minY {
			minY = position.Y
		} else if position.Y > maxY {
			maxY = position.Y
		}
	}

	// Add the node radius.
	maxX += float64(radius + borderWidth)
	maxY += float64(radius + borderWidth)
	minX -= float64(radius + borderWidth)
	minY -= float64(radius + borderWidth)

	inputWidth := maxX - minX

	scale := scaleWidth / inputWidth

	scaledCoordinates := make(map[string]tsne.Coordinate, len(coordinates))
	for gene, position := range coordinates {
		scaledCoordinates[gene] = tsne.Coordinate{
			X: scale * (position.X - minX),
			Y: scale * (position.Y - minY),
		}
	}

	return scaledCoordinates, scale * (maxY - minY)
}
