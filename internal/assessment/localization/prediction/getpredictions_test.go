package prediction

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var predictionNMFText = `gene	rank	score
AAAS	1	0.3326865	
AAK1	1	0.1396383	
AAR2	2	0.034698	
AARS2	2	0.2251458	
AASDH	1	0.09391108
`

var predictionSAFEText = `Node label	Node label ORF	Domain (predominant)	Neighborhood score [max=1, min=0] (predominant)	Total number of enriched domains	Number of enriched attributes per domain
VAMP3	VAMP3	2	0.600	1	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,0,0,0
SNAP29	SNAP29	1	0.263	0	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
CDCA3	CDCA3	2	1.000	1	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,61,0,0,0,0
`

var _ = Describe("Read predicted GO IDs", func() {
	It("should return GO IDs for every NMF prey prediction", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/prediction.txt",
			[]byte(predictionNMFText),
			0444,
		)
		options := parameters{
			predictions:       "test/prediction.txt",
			predictionSummary: "test/prediction-summary.txt",
			predictionType:    "nmf",
		}

		expected := map[string]int{
			"AAAS":  1,
			"AAK1":  1,
			"AAR2":  2,
			"AARS2": 2,
			"AASDH": 1,
		}
		Expect(getPredictions(options)).To(Equal(expected))
	})

	It("should return GO IDs for every NMF prey prediction", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/prediction.txt",
			[]byte(predictionSAFEText),
			0444,
		)
		options := parameters{
			predictions:       "test/prediction.txt",
			predictionSummary: "test/prediction-summary.txt",
			predictionType:    "safe",
		}

		expected := map[string]int{
			"VAMP3":  2,
			"SNAP29": 1,
			"CDCA3":  2,
		}
		Expect(getPredictions(options)).To(Equal(expected))
	})
})
