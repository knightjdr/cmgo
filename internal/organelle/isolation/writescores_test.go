package isolation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

var _ = Describe("Write score", func() {
	It("should write data to tsv file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		scores := &isolationScores{
			1: &isolationScore{
				edgesOutside:       3,
				edgesWithin:        1,
				isolation:          0.25,
				nodesOutside:       []int{1, 2, 2},
				sharedCompartments: []int{0, 0, 3},
			},
			2: &isolationScore{
				edgesOutside:       3,
				edgesWithin:        0,
				isolation:          0,
				nodesOutside:       []int{0, 0, 3},
				sharedCompartments: []int{0, 3, 0},
			},
		}
		summary := localization.Summary{
			1: localization.CompartmentSummary{
				DisplayTerms: []string{"compartmentA"},
			},
			2: localization.CompartmentSummary{
				DisplayTerms: []string{"compartmentB1", "compartmentB2"},
			},
		}

		expected := "compartment\tname\tisolation score\tedges within\tedges outside\n" +
			"1\tcompartmentA\t0.250\t1\t3\n" +
			"2\tcompartmentB1, compartmentB2\t0.000\t0\t3\n"
		writeScores(scores, summary, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
