package heatmap

import (
	"testing"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var enrichmentText = `rank	term	matched	background_size	fold enrichment	pvalue	adj. pvalue	bhfdr	genes
1	PDZ	23	48	2	1.62980740014973e-13	4.27009538839229e-11	3.81679389312977e-05	AFDN,CASK,DLG1
1	FERM_C	12	17	4	2.1798629399977e-12	2.85562045139698e-10	7.63358778625954e-05	EPB41,EPB41L1,EPB41L2
1	FERM_N	12	18	8	6.23068998013292e-12	5.44146924931608e-10	0.000114503816793893	EPB41,EPB41L1,EPB41L2
2	KRAB	24	28	2	2.32442462867683e-22	1.02972011050383e-19	2.25733634311512e-05	POGK,RBAK,ZFP1
2	zf-C2H2	30	65	4	2.99760216648792e-15	6.63968879877075e-13	4.51467268623025e-05	CTCF,MAZ,PRDM15
2	Bromodomain	13	16	6.259	1.29285132936678e-08	0.02	6.77200902934537e-05	ATAD2,BPTF,BRD1
2	PWWP	10	14	1	1.53538226513889e-08	1.70043585864132e-06	9.0293453724605e-05	BRD1,BRPF3,DNMT3A,GLYR1
`

func TestReadEnrichment(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/enrichment.txt",
		[]byte(enrichmentText),
		0444,
	)

	wanted := map[string]map[string]float64{
		"1": map[string]float64{
			"PDZ":    1,
			"FERM_C": 2,
			"FERM_N": 3,
		},
		"2": map[string]float64{
			"KRAB":    1,
			"zf-C2H2": 2,
			"PWWP":    0,
		},
	}
	enrichment := readEnrichment("test/enrichment.txt", 0.01)
	assert.Equal(t, wanted, enrichment, "Should read and filter by p-Value an enrichment file")
}
