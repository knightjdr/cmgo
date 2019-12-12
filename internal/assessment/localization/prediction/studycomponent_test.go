package prediction

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var hpaText = `gene	terms	go_id
TSPAN6	Cytosol	GO:0005829
SCYL3	Microtubules;Nuclear bodies	GO:0015630;GO:0016604
C1orf112	Mitochondria	GO:0005739
FGR	Aggresome;Plasma membrane	GO:0016235;GO:0005886
CFH	Vesicles	GO:0043231
GENE	Unknown	
`

var _ = Describe("Read predictions from another study", func() {
	It("should read predictions", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/hpa.txt",
			[]byte(hpaText),
			0444,
		)

		expected := map[string][]string{
			"C1orf112": []string{"GO:0005739"},
			"CFH":      []string{"GO:0043231"},
			"FGR":      []string{"GO:0016235", "GO:0005886"},
			"SCYL3":    []string{"GO:0015630", "GO:0016604"},
			"TSPAN6":   []string{"GO:0005829"},
		}
		Expect(readFractionationPredictions("test/hpa.txt", 2)).To(Equal(expected))
	})
})
