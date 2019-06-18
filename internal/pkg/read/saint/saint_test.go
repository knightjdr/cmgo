package saint

import (
	"testing"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var saintText = `Bait	Prey	PreyGene	Spec	SpecSum	AvgSpec	NumReplicates	ctrlCounts	AvgP	MaxP	TopoAvgP	TopoMaxP	SaintScore	logOddsScore	FoldChange	BFDR	boosted_by	UniqueSpec	UniqueSpecSum	UniqueAvgSpec	PreySequenceLength	UniProtID
AARS2	NP_000009.1	ACADVL	3|6	9	4.50	2	0|0|0|0|0|0	1.00	1.00	1.00	1.00	1.00	5.35	45.00	0.00	NP_057665.2|NP_789845.1|	3|6	9	4.5	655	NP_000009.1
AARS2	NP_000010.1	ACAT1	6|8	14	7.00	2	0|0|0|0|0|0	1.00	1.00	1.00	1.00	1.00	10.41	70.00	0.00	NP_000099.2|NP_001243439.1|	6|8	14	7	427	NP_000010.1
AARS2	NP_000091.1	CSTB	2|2	4	2.00	2	0|1|2|1|6|2	0.00	0.00	0.00	0.00	0.00	-11.72	0.49	0.54		2|2	4	2	98	NP_000091.1
AARS2	NP_000099.2	DLD	18|14	32	16.00	2	0|0|0|1|0|0	1.00	1.00	1.00	1.00	1.00	9.13	22.86	0.00	NP_001014763.1|NP_001135858.1|	18|14	32	16	509	NP_000099.2
AARS2	NP_000131.2	FECH	4|1	5	2.50	2	0|0|0|0|0|0	0.50	1.00	0.50	1.00	0.50	1.49	25.00	0.09	NP_001258625.1|	4|1	5	2.5	423	NP_000131.2
ABCC1	NP_000009.1	ACADVL	3|6	9	4.50	2	0|0|0|0|0|0	1.00	1.00	1.00	1.00	1.00	5.35	45.00	0.00	NP_057665.2|NP_789845.1|	3|6	9	4.5	655	NP_000009.1
`

func TestMapSaintLine(t *testing.T) {
	// TEST1: at least 21 elements
	line := []string{
		"AARS2",
		"NP_000009.1",
		"ACADVL",
		"3|6",
		"9",
		"4.50",
		"2",
		"0|0|0|0|0|0|0",
		"1.00",
		"1.00",
		"1.00",
		"1.00",
		"1.00",
		"5.35",
		"45.00",
		"0.00",
		"NP_057665.2|NP_789845.1|NP_001536.1|NP_055177.2|",
		"3|6",
		"9",
		"4.5",
		"655",
		"NP_000009.1",
	}
	wanted := Row{
		Bait:               "AARS2",
		Prey:               "NP_000009.1",
		PreyGene:           "ACADVL",
		AvgSpec:            4.5,
		Control:            []float64{0, 0, 0, 0, 0, 0, 0},
		AvgP:               1,
		FoldChange:         45,
		FDR:                0,
		PreySequenceLength: 655,
		Spec:               []float64{3, 6},
	}
	assert.Equal(t, wanted, mapSaintLine(line), "Should map line from SAINT file to struct")

	// TEST2: less than 21 elements
	line = []string{
		"AARS2",
		"NP_000009.1",
		"ACADVL",
		"3|6",
		"9",
		"4.50",
		"2",
		"0|0|0|0|0|0|0",
		"1.00",
		"1.00",
		"1.00",
		"1.00",
		"1.00",
		"5.35",
		"45.00",
		"0.00",
	}
	wanted = Row{
		Bait:               "AARS2",
		Prey:               "NP_000009.1",
		PreyGene:           "ACADVL",
		AvgSpec:            4.5,
		Control:            []float64{0, 0, 0, 0, 0, 0, 0},
		AvgP:               1,
		FoldChange:         45,
		FDR:                0,
		PreySequenceLength: 0,
		Spec:               []float64{3, 6},
	}
	assert.Equal(t, wanted, mapSaintLine(line), "Should map line from SAINT file to struct with nil value for prey length")
}

func TestFilterBaits(t *testing.T) {
	rows := []Row{
		{Bait: "AARS2", Prey: "NP_000009.1", PreyGene: "ACADVL", AvgSpec: 4.5, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 45, FDR: 0, PreySequenceLength: 655, Spec: []float64{3, 6}},
		{Bait: "AARS2", Prey: "NP_000010.1", PreyGene: "ACAT1", AvgSpec: 7, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 70, FDR: 0, PreySequenceLength: 427, Spec: []float64{6, 8}},
		{Bait: "AARS2", Prey: "NP_000099.2", PreyGene: "DLD", AvgSpec: 16, Control: []float64{0, 0, 0, 1, 0, 0}, AvgP: 1, FoldChange: 22.86, FDR: 0, PreySequenceLength: 509, Spec: []float64{18, 14}},
		{Bait: "ABCC1", Prey: "NP_000009.1", PreyGene: "ACADVL", AvgSpec: 4.5, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 45, FDR: 0, PreySequenceLength: 655, Spec: []float64{3, 6}},
	}

	// TEST1: do not filter when the minimum bait value is 1
	wanted := rows
	assert.Equal(t, wanted, filterBaits(rows, 1), "Should return input rows when there is no bait filter")

	// TEST2: should filter by minBait number
	wanted = []Row{
		{Bait: "AARS2", Prey: "NP_000009.1", PreyGene: "ACADVL", AvgSpec: 4.5, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 45, FDR: 0, PreySequenceLength: 655, Spec: []float64{3, 6}},
		{Bait: "ABCC1", Prey: "NP_000009.1", PreyGene: "ACADVL", AvgSpec: 4.5, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 45, FDR: 0, PreySequenceLength: 655, Spec: []float64{3, 6}},
	}
	assert.Equal(t, wanted, filterBaits(rows, 2), "Should filter to remove preys seen with less baits than bait filter")
}

func TestRead(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/saint.txt",
		[]byte(saintText),
		0444,
	)

	// TEST1: only filter by FDR
	wanted := []Row{
		{Bait: "AARS2", Prey: "NP_000009.1", PreyGene: "ACADVL", AvgSpec: 4.5, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 45, FDR: 0, PreySequenceLength: 655, Spec: []float64{3, 6}},
		{Bait: "AARS2", Prey: "NP_000010.1", PreyGene: "ACAT1", AvgSpec: 7, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 70, FDR: 0, PreySequenceLength: 427, Spec: []float64{6, 8}},
		{Bait: "AARS2", Prey: "NP_000099.2", PreyGene: "DLD", AvgSpec: 16, Control: []float64{0, 0, 0, 1, 0, 0}, AvgP: 1, FoldChange: 22.86, FDR: 0, PreySequenceLength: 509, Spec: []float64{18, 14}},
		{Bait: "ABCC1", Prey: "NP_000009.1", PreyGene: "ACADVL", AvgSpec: 4.5, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 45, FDR: 0, PreySequenceLength: 655, Spec: []float64{3, 6}},
	}
	assert.Equal(t, wanted, Read("test/saint.txt", 0.01, 1), "Should read and filter SAINT file by FDR")

	// TEST2: filter by FDR and minBait number
	wanted = []Row{
		{Bait: "AARS2", Prey: "NP_000009.1", PreyGene: "ACADVL", AvgSpec: 4.5, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 45, FDR: 0, PreySequenceLength: 655, Spec: []float64{3, 6}},
		{Bait: "ABCC1", Prey: "NP_000009.1", PreyGene: "ACADVL", AvgSpec: 4.5, Control: []float64{0, 0, 0, 0, 0, 0}, AvgP: 1, FoldChange: 45, FDR: 0, PreySequenceLength: 655, Spec: []float64{3, 6}},
	}
	assert.Equal(t, wanted, Read("test/saint.txt", 0.01, 2), "Should read and filter SAINT file by FDR and minimum bait")
}
