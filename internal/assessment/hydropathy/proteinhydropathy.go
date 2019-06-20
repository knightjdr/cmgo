package hydropathy

import "strings"

var aaScale = map[string]float64{
	"A": 1.8,
	"R": -4.5,
	"N": -3.5,
	"D": -3.5,
	"C": 2.5,
	"Q": -3.5,
	"E": -3.5,
	"G": -0.4,
	"H": -3.2,
	"I": 4.5,
	"L": 3.8,
	"K": -3.9,
	"M": 1.9,
	"F": 2.8,
	"P": -1.6,
	"S": -0.8,
	"T": -0.7,
	"W": -0.9,
	"Y": -1.3,
	"V": 4.2,
}

func proteinHydropathy(proteins []string, sequences map[string]string) []float64 {
	scores := make([]float64, 0)
	proteinIndex := 0
	for _, protein := range proteins {
		if _, ok := sequences[protein]; ok {
			aaSequence := strings.Split(sequences[protein], "")
			length := 0
			score := 0.0
			for _, aa := range aaSequence {
				if _, ok := aaScale[aa]; ok {
					score += aaScale[aa]
					length++
				}
			}
			scores = append(scores, score/float64(length))
			proteinIndex++
		}
	}
	return scores
}
