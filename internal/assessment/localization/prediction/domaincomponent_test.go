package prediction

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var compartmentDomainsText = `rank	term	matched	background_size	fold enrichment	pvalue	adj. pvalue	bhfdr	genes
1	PDZ	23	48	8.906	1.62980740014973e-13	4.27009538839229e-11	3.81679389312977e-05	AFDN,CASK
1	FERM_C	12	17	13.120	2.1798629399977e-12	2.85562045139698e-10	7.63358778625954e-05	EPB41,EPB41L5
2	KRAB	24	28	9.905	2.32442462867683e-22	1.02972011050383e-19	2.25733634311512e-05	POGK,RBAK
3	TRAM_LAG1_CLN8	8	9	30.815	3.37809753356502e-12	4.45908874430583e-10	7.57575757575758e-05	CERS1,CERS4
`

var geneDomainsText = `#<seq id> <alignment start> <alignment end> <envelope start> <envelope end> <hmm acc> <hmm name> <type> <hmm start> <hmm end> <hmm length> <bit score> <E-value> <clan>
D6RHZ6	163	285	163	299	PF00107	ADH_zinc_N	PfamLive::Result::SequenceOntology=HASH(0x89baa50)	1	117	130	91.50	1e-22	CL0063
D6RHZ6	2	120	1	121	PF08240	ADH_N	PfamLive::Result::SequenceOntology=HASH(0xf623a48)	9	108	109	83.60	2e-20	CL0296
O95863	153	176	153	177	PF13912	zf-C2H2_6	PfamLive::Result::SequenceOntology=HASH(0x108198b8)	1	24	27	20.50	0.9	CL0361
O95863	208	230	208	230	PF00096	zf-C2H2	PfamLive::Result::SequenceOntology=HASH(0x89b8a80)	1	23	23	26.70	0.013	CL0361
O95863	181	202	180	202	PF00096	zf-C2H2	PfamLive::Result::SequenceOntology=HASH(0x89b8a80)	2	23	23	23.40	0.14	CL0361
`

var _ = Describe("Read compartment domains", func() {
	It("should return domains for each compartment", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/compartment-domains.txt",
			[]byte(compartmentDomainsText),
			0444,
		)

		expected := map[int][]string{
			1: []string{"PDZ", "FERM_C"},
			2: []string{"KRAB"},
			3: []string{"TRAM_LAG1_CLN8"},
		}
		Expect(readCompartmentDomains("test/compartment-domains.txt")).To(Equal(expected))
	})
})

var _ = Describe("Read gene domains", func() {
	It("should return domains for each gene", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/gene-domains.txt",
			[]byte(geneDomainsText),
			0444,
		)
		geneToUniProt := map[string]string{
			"prey1": "D6RHZ6",
			"prey2": "",
			"prey3": "O95863",
		}
		predictions := map[string]int{
			"prey1": 1,
			"prey2": 2,
			"prey3": 1,
		}
		inputFiles := fileContent{
			geneToUniProt: geneToUniProt,
			predictions:   predictions,
		}

		expected := map[string][]string{
			"prey1": []string{"ADH_N", "ADH_zinc_N"},
			"prey2": []string{},
			"prey3": []string{"zf-C2H2", "zf-C2H2", "zf-C2H2_6"},
		}
		Expect(readGeneDomains("test/gene-domains.txt", inputFiles)).To(Equal(expected))
	})
})
