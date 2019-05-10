package nmfsafe

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestSummarize(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)

	genes := map[string]*localizationInfo{
		"a": &localizationInfo{
			Concordant: true,
			Domain:     1,
			NMFids:     []string{"GO:1", "GO:2"},
			NMFinSAFE:  false,
			NMFknown:   true,
			Rank:       1,
			SAFEids:    []string{"GO:3"},
			SAFEinNMF:  true,
			SAFEknown:  false,
		},
		"b": &localizationInfo{
			Concordant: true,
			Domain:     2,
			NMFids:     []string{"GO:4"},
			NMFinSAFE:  true,
			NMFknown:   true,
			Rank:       3,
			SAFEids:    []string{"GO:4", "GO:5"},
			SAFEinNMF:  true,
			SAFEknown:  true,
		},
		"c": &localizationInfo{
			Concordant: false,
			NMFids:     []string{"GO:4"},
			NMFinSAFE:  false,
			NMFknown:   false,
			Rank:       3,
			SAFEids:    []string{},
			SAFEinNMF:  false,
			SAFEknown:  false,
		},
	}
	wanted := "Total genes:\t3\n" +
		"Genes with NMF and SAFE assignment:\t2\n" +
		"Concordant genes:\t2\t100.00\n\n" +
		"analysis\ttotal genes\tassigned term\t% assigned\tknown\t% known\n" +
		"NMF\t3\t3\t100.00\t2\t66.67\n" +
		"SAFE\t2\t2\t100.00\t1\t50.00\n"
	summarize(genes, "test/out.txt")
	bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
	assert.Equal(t, wanted, string(bytes), "Should write summary to file")
}
