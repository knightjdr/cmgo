package rbo

import customMath "github.com/knightjdr/cmgo/pkg/math"

func truncate(s, t []string, userK int) ([]string, []string) {
	if userK == 0 {
		return s, t
	}

	var maxK int
	if len(s) > len(t) {
		maxK = len(s)
	} else {
		maxK = len(t)
	}

	var k int
	if userK > maxK {
		k = maxK
	} else {
		k = userK
	}

	return s[:customMath.MinInt(len(s), k)], t[:customMath.MinInt(len(t), k)]
}
