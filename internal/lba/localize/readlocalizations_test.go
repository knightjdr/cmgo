package localize

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var localizationText = `GO:0005654	nucleoplasm
GO:0005694	chromosome
GO:0005730	nucleolus
`

var _ = Describe("Read localization file", func() {
	It("should read file as a map of GO IDs to names", func() {
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
		expected := map[string]string{
			"GO:0005654": "nucleoplasm",
			"GO:0005694": "chromosome",
			"GO:0005730": "nucleolus",
		}
		Expect(readLocalizations("test/localization.txt")).To(Equal(expected))
	})
})
