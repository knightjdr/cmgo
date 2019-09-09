package dbgenes

import (
	"sort"

	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var databaseText = `#tax_id	GeneID	Symbol	LocusTag	Synonyms	dbXrefs	chromosome	map_location	description	type_of_gene	Symbol_from_nomenclature_authority	Full_name_from_nomenclature_authority	Nomenclature_status	Other_designations	Modification_date	Feature_type
192	36107229	AMK58_RS00870	AMK58_RS00870	AMK58_00870	-	-	-	2-polyprenylphenol 6-hydroxylase	protein-coding	-	-	-	-	20180309	-
9606	1	A1BG	-	A1B|ABG|GAB|HYST2477	MIM:138670|HGNC:HGNC:5|Ensembl:ENSG00000121410	19	19q13.43	alpha-1-B glycoprotein	protein-coding	A1BG	alpha-1-B glycoprotein	O	alpha-1B-glycoprotein|HEL-S-163pA|epididymis secretory sperm binding protein Li 163pA	20190617	-
9606	2	A2M	-	A2MD|CPAMD5|FWP007|S863-7	MIM:103950|HGNC:HGNC:7|Ensembl:ENSG00000175899	12	12p13.31	alpha-2-macroglobulin	protein-coding	A2M	alpha-2-macroglobulin	O	alpha-2-macroglobulin|C3 and PZP-like alpha-2-macroglobulin domain-containing protein 5|alpha-2-M	20190817	-
`

var _ = Describe("ReadList", func() {
	It("should read human gene names from list", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/gene_info.txt",
			[]byte(databaseText),
			0444,
		)

		expected := []string{"A1BG", "A2M"}
		results := readList("test/gene_info.txt")
		sort.Strings(results)
		Expect(results).To(Equal(expected), "should read official gene names")
	})
})
