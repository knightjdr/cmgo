package rbo

import (
	customMath "github.com/knightjdr/cmgo/pkg/math"
	"github.com/knightjdr/cmgo/pkg/slice"
)

// I is the intersection between lists S and T
func I(S, T []string, d int) []string {
	endS := customMath.MinInt(len(S), d)
	endT := customMath.MinInt(len(T), d)
	return slice.Intersect(S[:endS], T[:endT])
}
