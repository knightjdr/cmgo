package localization

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var localizationText = `Symbol	Merged localizations	Merged GO
A1BG
A2M	extracellular region	GO:0005576
A2ML1	plasma membrane;nucleus ;mitochondrion;cytoplasm	GO:0005886; GO:0005634 ;GO:0005739;GO:0005737
A4GALT	cytoplasm	GO:0005737
`

var _ = Describe("Read prey localization file", func() {
	It("should read file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/localization.txt",
			[]byte(localizationText),
			0444,
		)

		expected := map[string]map[string]string{
			"A1BG": map[string]string{},
			"A2M": map[string]string{
				"GO:0005576": "extracellular region",
			},
			"A2ML1": map[string]string{
				"GO:0005886": "plasma membrane",
				"GO:0005634": "nucleus",
				"GO:0005739": "mitochondrion",
				"GO:0005737": "cytoplasm",
			},
			"A4GALT": map[string]string{
				"GO:0005737": "cytoplasm",
			},
		}
		Expect(Prey("test/localization.txt")).To(Equal(expected))
	})
})
