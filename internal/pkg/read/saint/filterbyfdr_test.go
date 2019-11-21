package saint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterByFDR(t *testing.T) {
	saint := &SAINT{
		Row{Bait: "AARS2", Prey: "NP_000009.1", PreyGene: "ACADVL", AvgSpec: 4.5, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 45, FDR: 0, PreySequenceLength: 655, Spec: []float64{3, 6}},
		Row{Bait: "AARS2", Prey: "NP_000010.1", PreyGene: "ACAT1", AvgSpec: 7, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 70, FDR: 0.05, PreySequenceLength: 427, Spec: []float64{6, 8}},
		Row{Bait: "AARS2", Prey: "NP_000099.2", PreyGene: "DLD", AvgSpec: 16, Control: []float64{0, 0, 0, 1, 0, 0}, AvgP: 1, FoldChange: 22.86, FDR: 0.02, PreySequenceLength: 509, Spec: []float64{18, 14}},
		Row{Bait: "ABCC1", Prey: "NP_000009.1", PreyGene: "ACADVL", AvgSpec: 4.5, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 45, FDR: 0.03, PreySequenceLength: 655, Spec: []float64{3, 6}},
	}

	wanted := &SAINT{
		Row{Bait: "AARS2", Prey: "NP_000009.1", PreyGene: "ACADVL", AvgSpec: 4.5, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 45, FDR: 0, PreySequenceLength: 655, Spec: []float64{3, 6}},
		Row{Bait: "AARS2", Prey: "NP_000099.2", PreyGene: "DLD", AvgSpec: 16, Control: []float64{0, 0, 0, 1, 0, 0}, AvgP: 1, FoldChange: 22.86, FDR: 0.02, PreySequenceLength: 509, Spec: []float64{18, 14}},
	}
	saint.FilterByFDR(0.02)
	assert.Equal(t, wanted, saint, "Should filter rows not passing FDR")
}
