package prediction

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var baitExpected = `id	bait	term	go_id
1	bait1	aaaa	GO:1111
2	bait2	aaaa	GO:1111
3	bait3	bbbb	GO:2222
`

var predictionSummaryText = `rank	term	displayname	go	synonyms	ic
1	[compartment1]	[compartment1]	[GO:1111, GO:2222]	[]	[1.166]
2	[compartment2]	[compartment2]	[GO:2222]	[]	[1.256]
`

var _ = Describe("Read bait localizations", func() {
	It("should return bait localizations as compartment IDs", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/bait-expected.txt",
			[]byte(baitExpected),
			0444,
		)
		afero.WriteFile(
			fs.Instance,
			"test/summary.txt",
			[]byte(predictionSummaryText),
			0444,
		)
		goHierarchy := &geneontology.GOhierarchy{
			"CC": map[string]*geneontology.GOterm{
				"GO:1111": &geneontology.GOterm{
					Children: []string{"GO:1111_1", "GO:1111_2"},
				},
				"GO:2222": &geneontology.GOterm{
					Children: []string{"GO:2222_1", "GO:2222_2"},
				},
			},
		}
		summary := localization.Summary{
			1: localization.CompartmentSummary{
				GOid: []string{"GO:1111", "GO:2222"},
			},
			2: localization.CompartmentSummary{
				GOid: []string{"GO:2222"},
			},
		}
		inputFiles := fileContent{
			goHierarchy:       goHierarchy,
			predictionSummary: summary,
		}

		expected := baitInformation{
			compartmentCounts: map[int]int{
				1: 3,
				2: 1,
			},
			localizations: map[string][]int{
				"bait1": []int{1},
				"bait2": []int{1},
				"bait3": []int{1, 2},
			},
		}
		Expect(getBaitLocalizationsAsCompartments("test/bait-expected.txt", inputFiles)).To(Equal(expected))
	})
})
