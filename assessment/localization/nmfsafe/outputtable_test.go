package nmfsafe

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestOutputTable(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)

	genes := map[string]*localizationInfo{
		"a": &localizationInfo{
			Concordant: true,
			Domain:     1,
			NMFinSAFE:  false,
			NMFknown:   true,
			NMFterms:   []string{"GO:1", "GO:2"},
			Rank:       2,
			SAFEinNMF:  true,
			SAFEknown:  false,
			SAFEterms:  []string{"GO:3"},
		},
		"b": &localizationInfo{
			Concordant: true,
			Domain:     3,
			NMFinSAFE:  true,
			NMFknown:   true,
			NMFterms:   []string{"GO:4"},
			Rank:       5,
			SAFEinNMF:  true,
			SAFEknown:  true,
			SAFEterms:  []string{"GO:4", "GO:5"},
		},
	}
	wanted := "gene\trank\tNMF term(s)\tNMF known?\tdomain\tSAFE term(s)\tSAFE known?\tNMF in SAFE?\tSAFE in NMF?\tconcordant\n" +
		"a\t2\tGO:1, GO:2\ttrue\t1\tGO:3\tfalse\tfalse\ttrue\ttrue\n" +
		"b\t5\tGO:4\ttrue\t3\tGO:4, GO:5\ttrue\ttrue\ttrue\ttrue\n"
	outputTable(genes, "test/out.txt")
	bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
	assert.Equal(t, wanted, string(bytes), "Should write gene results to tabular file")
}
