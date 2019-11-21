package interactions

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var biogridText = `#BioGRID Interaction ID	Entrez Gene Interactor A	Entrez Gene Interactor B	BioGRID ID Interactor A	BioGRID ID Interactor B	Systematic Name Interactor A	Systematic Name Interactor B	Official Symbol Interactor A	Official Symbol Interactor B
103	6416	2318	112315	108607	-	-	MAP2K4	FLNC
183	90	2339	106605	108625	-	-	ACVR1	FNTA
612	377	23647	106872	117174	-	-	ARF3	ARFIP2
617	377	27236	106872	118084	-	-	ARF3	ARFIP1
618	23647	377	117174	106872	-	-	ARFIP2	ARF3
619	27236	377	118084	106872	-	-	ARFIP1	ARF3
663	54464	226	119970	106728	-	-	XRN1	ALDOA
`

var intactText = `#ID(s) interactor A	ID(s) interactor B	Alt. ID(s) interactor A	Alt. ID(s) interactor B	Alias(es) interactor A	Alias(es) interactor B	Interaction detection method(s)	Publication 1st author(s)	Publication Identifier(s)	Taxid interactor A
uniprotkb:P49418	uniprotkb:O43426	intact:EBI-7121510	intact:EBI-2821539	psi-mi:amph_human(display_long)|uniprotkb:AMPH(gene name)|psi-mi:AMPH(display_short)|uniprotkb:AMPH1(gene name synonym)	psi-mi:synj1_human(display_long)|uniprotkb:SYNJ1(gene name)|psi-mi:SYNJ1(display_short)|uniprotkb:KIAA0910(gene name synonym)|uniprotkb:Synaptic inositol 1,4,5-trisphosphate 5-phosphatase 1(gene name synonym)	psi-mi:"MI:0084"(phage display)	Cestra et al. (1999)	pubmed:10542231|mint:MINT-5211933	taxid:9606(human)|taxid:9606(Homo sapiens)	taxid:9606(human)|taxid:9606(Homo sapiens)
uniprotkb:Q99961	uniprotkb:Q05193	intact:EBI-697911	intact:EBI-713135	psi-mi:sh3g1_human(display_long)|uniprotkb:SH3GL1(gene name)|psi-mi:SH3GL1(display_short)|uniprotkb:CNSA1(gene name synonym)|uniprotkb:SH3D2B(gene name synonym)	psi-mi:dyn1_human(display_long)|uniprotkb:DNM1(gene name)|psi-mi:DNM1(display_short)|uniprotkb:DNM(gene name synonym)	psi-mi:"MI:0081"(peptide array)	Cestra et al. (1999)	pubmed:10542231|mint:MINT-5211933	taxid:10090(mouse)|taxid:10090(Mus musculus)	taxid:10090(mouse)|taxid:10090(Mus musculus)
`

var _ = Describe("Read interaction file", func() {
	It("should read interactors from biogrid data", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/biogrid.txt",
			[]byte(biogridText),
			0444,
		)

		expected := map[string][]string{
			"ACVR1":  []string{"FNTA"},
			"ALDOA":  []string{"XRN1"},
			"ARF3":   []string{"ARFIP1", "ARFIP2"},
			"ARFIP1": []string{"ARF3"},
			"ARFIP2": []string{"ARF3"},
			"FLNC":   []string{"MAP2K4"},
			"FNTA":   []string{"ACVR1"},
			"MAP2K4": []string{"FLNC"},
			"XRN1":   []string{"ALDOA"},
		}
		Expect(readFile("test/biogrid.txt", "", parseBiogridLine)).To(Equal(expected))
	})

	It("should read interactors from intact data", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/intact.txt",
			[]byte(intactText),
			0444,
		)

		expected := map[string][]string{
			"AMPH":  []string{"SYNJ1"},
			"SYNJ1": []string{"AMPH"},
		}
		Expect(readFile("test/intact.txt", "9606", parseIntactLine)).To(Equal(expected))
	})
})
