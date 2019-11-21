package saint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthNormalizeSpectralCounts(t *testing.T) {
	saint := &SAINT{
		Row{Bait: "AARS2", AvgSpec: 10, Control: []float64{2, 2}, PreySequenceLength: 100},
		Row{Bait: "AARS2", AvgSpec: 20, Control: []float64{6, 8}, PreySequenceLength: 50},
		Row{Bait: "AARS2", AvgSpec: 50, Control: []float64{0, 0}, PreySequenceLength: 150},
		Row{Bait: "ABCC1", AvgSpec: 100, Control: []float64{0, 50}, PreySequenceLength: 200},
		Row{Bait: "ABCC1", AvgSpec: 80, Control: []float64{4, 4}},
		Row{Bait: "ABCC1", AvgSpec: 20, Control: []float64{0, 0}, PreySequenceLength: 500},
	}
	wanted := &SAINT{
		Row{Bait: "AARS2", AvgSpec: 10, Control: []float64{2, 2}, NormalizedSpec: 8, PreySequenceLength: 100},
		Row{Bait: "AARS2", AvgSpec: 20, Control: []float64{6, 8}, NormalizedSpec: 26, PreySequenceLength: 50},
		Row{Bait: "AARS2", AvgSpec: 50, Control: []float64{0, 0}, NormalizedSpec: 33.33, PreySequenceLength: 150},
		Row{Bait: "ABCC1", AvgSpec: 100, Control: []float64{0, 50}, NormalizedSpec: 131.25, PreySequenceLength: 200},
		Row{Bait: "ABCC1", AvgSpec: 80, Control: []float64{4, 4}, NormalizedSpec: 76},
		Row{Bait: "ABCC1", AvgSpec: 20, Control: []float64{0, 0}, NormalizedSpec: 14, PreySequenceLength: 500},
	}

	saint.LengthNormalizeSpectralCounts()
	assert.Equal(t, wanted, saint, "Should add control subtracted and length adjusted spectral counts to SAINT variable")
}
