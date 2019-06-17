package rbo

import (
	"math"

	customMath "github.com/knightjdr/cmgo/pkg/math"
)

// RBOext calculates the rank biased overlap to depth k, based on formula 32
// from "A Similarity Measure for Indefinite Rankings" (Webber et al.).
// k should not be greater than max(len(s), len(t)), and will be set to this
// minimum if it is. Set k to 0 to compare full length of lists.
// p is the persitence (weighting).
func RBOext(Sinput, Tinput []string, p float64, userK int) float64 {
	S, T := truncate(Sinput, Tinput, userK)
	l := customMath.MaxInt(len(S), len(T))
	s := customMath.MinInt(len(S), len(T))
	if s == 0 {
		return 0
	}

	Xl := X(I(S, T, l))
	Xs := X(I(S, T, s))

	// Calculate overlap until l.
	simL := float64(0)
	for d := 1; d <= l; d++ {
		Xd := X(I(S, T, d))
		simL += (float64(Xd) / float64(d)) * math.Pow(p, float64(d))
	}

	// Calculate overlap after s.
	simAfterS := float64(0)
	for d := s + 1; d <= l; d++ {
		simAfterS += ((float64(Xs) * float64(d-s)) / float64(s*d)) * math.Pow(p, float64(d))
	}

	rbo := (((1 - p) / p) * (simL + simAfterS)) + (((float64(Xl-Xs) / float64(l)) + (float64(Xs) / float64(s))) * math.Pow(p, float64(l)))
	if rbo > 1 {
		rbo = 1
	}
	return rbo
}
