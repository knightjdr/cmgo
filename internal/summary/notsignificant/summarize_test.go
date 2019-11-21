package notsignificant

import (
	"testing"

	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	"github.com/stretchr/testify/assert"
)

func TestSummarize(t *testing.T) {
	data := &saint.SAINT{
		{Bait: "a", PreyGene: "prey-1", Spec: []float64{10, 5}, Control: []float64{0, 4, 8}, FDR: 0.01},
		{Bait: "a", PreyGene: "prey-2", Spec: []float64{15, 13}, Control: []float64{5, 10, 30}, FDR: 0},
		{Bait: "a", PreyGene: "prey-3", Spec: []float64{4, 2}, Control: []float64{20, 30, 25}, FDR: 0.05},
		{Bait: "b", PreyGene: "prey-1", Spec: []float64{3.7, 2}, Control: []float64{0, 4, 8}, FDR: 0.03},
		{Bait: "b", PreyGene: "prey-3", Spec: []float64{5, 5}, Control: []float64{20, 30, 25}, FDR: 0.01},
		{Bait: "c", PreyGene: "prey-1", Spec: []float64{47, 30}, Control: []float64{0, 4, 8}, FDR: 0.09},
		{Bait: "c", PreyGene: "prey-2", Spec: []float64{12, 11}, Control: []float64{5, 10, 30}, FDR: 0.1},
		{Bait: "c", PreyGene: "prey-4", Spec: []float64{35.5, 32}, Control: []float64{0, 0, 0}, FDR: 0},
	}
	wanted := map[string]*preySummary{
		"prey-1": &preySummary{baits: []string{"a", "b", "c"}, bestFDR: 0.01, ctrlAvg: 4, maxSpec: 47},
		"prey-2": &preySummary{baits: []string{"a", "c"}, bestFDR: 0, ctrlAvg: 15, maxSpec: 15},
		"prey-3": &preySummary{baits: []string{"a", "b"}, bestFDR: 0.01, ctrlAvg: 25, maxSpec: 5},
		"prey-4": &preySummary{baits: []string{"c"}, bestFDR: 0, ctrlAvg: 0, maxSpec: 35.5},
	}
	assert.Equal(t, wanted, summarize(data), "Should summarize SAINT data by prey")
}
