package moonlighting

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write moonlighting scores", func() {
	It("should write moonlighting scores to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		options := outputOptions{
			localization: localization.Summary{
				1: localization.CompartmentSummary{
					DisplayTerms: []string{"compartment 1"},
				},
				2: localization.CompartmentSummary{
					DisplayTerms: []string{"compartment 2"},
				},
				3: localization.CompartmentSummary{
					DisplayTerms: []string{"compartment 3"},
				},
				4: localization.CompartmentSummary{
					DisplayTerms: []string{"compartment 4"},
				},
				5: localization.CompartmentSummary{
					DisplayTerms: []string{"compartment 5"},
				},
			},
			minRankValue: 0.15,
			outfile:      "test/out.txt",
			preyNames:    []string{"geneA", "geneB", "geneC"},
		}
		scores := moonScores{
			&preyInfo{
				MoonlightingScore: 0.5,
				PrimaryRank:       1,
				PrimaryScore:      0.5,
				SecondaryRank:     2,
				SecondaryScore:    0.25,
			},
			&preyInfo{
				MoonlightingScore: 0.75,
				PrimaryRank:       0,
				PrimaryScore:      1,
				SecondaryRank:     1,
				SecondaryScore:    0.75,
			},
			&preyInfo{
				MoonlightingScore: 0.5,
				PrimaryRank:       4,
				PrimaryScore:      0.2,
				SecondaryRank:     0,
				SecondaryScore:    0.1,
			},
		}

		expected := "prey\t1st rank name\t2nd rank name\tmoonlighting score\t1st rank\t1st score\t2nd rank\t2nd score\n" +
			"geneA\tcompartment 2\tcompartment 3\t0.500\t2\t0.50000\t3\t0.25000\n" +
			"geneB\tcompartment 1\tcompartment 2\t0.750\t1\t1.00000\t2\t0.75000\n" +
			"geneC\tcompartment 5\t\t\t5\t0.20000\t\t\n"

		writeMoonlightingScores(scores, options)
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
