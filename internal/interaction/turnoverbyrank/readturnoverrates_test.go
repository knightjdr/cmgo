package turnoverbyrank

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var turnoverText = `ProteinGroup.id	UniProt identifier(s)	Gene name(s)	Protein name(s)	Potential contaminant	Number of proteins in group	Unique peptides	Sequence coverage [%]	Valid values for K	Cell culture replicate values for K	Protein sequence length	Molecular weight [kDa]	K [h-1]	k [h-1]	T50% [h]	T1/2 [h]
3451	Q13685	AAMP	Angio-associated migratory cell protein		1	13	41.2	8	4	434	46.75	0.089305319	0.065377063	7.761544199	10.60229913
4121	Q5JTZ9	AARS2	Alanine--tRNA ligase;mitochondrial		1	16	22.4	8	4	985	107.34	0.024570714	0.016213149	28.21029836	42.75216253
7546	Q9NY61	AATF	Protein AATF		1	10	25.9	8	4	560	63.132	0.060043984	0.035966671	11.54399052	19.27193058
2753	P61221	ABCE1	ATP-binding cassette sub-family E member 1		1	24	41.1	8	4	599	67.314	0.036091827	0.009979289	19.20510076	69.45857683
2234	P42765	ACAA2;ACAA3	3-ketoacyl-CoA thiolase;mitochondrial		1	18	61.7	8	4	397	41.924	0.026792723	0.00656	25.87072555	105.66268
`

var _ = Describe("Read turnover rate file", func() {
	It("should read gene names with turnover rates", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/turnover.txt",
			[]byte(turnoverText),
			0444,
		)

		expected := map[string]float64{
			"AAMP":  10.60229913,
			"AARS2": 42.75216253,
			"AATF":  19.27193058,
			"ABCE1": 69.45857683,
			"ACAA2": 105.66268,
			"ACAA3": 105.66268,
		}
		Expect(readTurnoverRates("test/turnover.txt")).To(Equal(expected))
	})
})
