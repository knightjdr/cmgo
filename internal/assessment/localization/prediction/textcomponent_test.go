package prediction

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var textAnnotationsText = `ENSP00000255087	MTL5	GO:0042645	Mitochondrial nucleoid	1.010	0.505	
ENSP00000255087	MTL5	GO:0000229	Cytoplasmic chromosome	1.009	0.504	
ENSP00000255087	MTL5	GO:0045495	Pole plasm	1.004	0.502	
ENSP00000255108	DPH2	GO:0005840	Ribosome	3.332	1.666	
ENSP00000255108	DPH2	GO:1990904	Ribonucleoprotein complex	3.077	1.538
`

var _ = Describe("Read text annotations", func() {
	It("should return GO IDs and text annotation score", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/text-annotations.txt",
			[]byte(textAnnotationsText),
			0444,
		)

		expected := map[string]map[string]float64{
			"DPH2": map[string]float64{
				"GO:0005840": 1.666,
				"GO:1990904": 1.538,
			},
			"MTL5": map[string]float64{
				"GO:0000229": 0.504,
				"GO:0042645": 0.505,
				"GO:0045495": 0.502,
			},
		}
		Expect(readTextAnnotations("test/text-annotations.txt")).To(Equal(expected))
	})
})
