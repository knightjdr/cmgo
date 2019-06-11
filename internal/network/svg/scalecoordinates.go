package svg

import "math"

func scaleCoordinates(coordinates map[string]coordinate, scaleWidth float64, radius, borderWidth int) (map[string]coordinate, float64) {
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

	scaledCoordinates := make(map[string]coordinate, len(coordinates))
	for gene, position := range coordinates {
		scaledCoordinates[gene] = coordinate{
			X: scale * (position.X - minX),
			Y: scale * (position.Y - minY),
		}
	}

	return scaledCoordinates, scale * (maxY - minY)
}
