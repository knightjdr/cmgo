package prediction

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write prey scores", func() {
	It("should write all preys with prediction scores", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		predictions := map[string]int{
			"prey1": 1,
			"prey2": 2,
		}
		predictionSummary := localization.Summary{
			1: localization.CompartmentSummary{
				GOid:    []string{"GO:111111"},
				GOterms: []string{"Term 1"},
			},
			2: localization.CompartmentSummary{
				GOid:    []string{"GO:222222"},
				GOterms: []string{"Term 2"},
			},
		}
		scores := preyScore{
			Bait: &preyBaitScore{
				"prey1": &baitScoreComponents{
					baits: []string{"bait1"},
					score: 0.11111,
				},
				"prey2": &baitScoreComponents{
					baits: []string{"bait2", "bait3"},
					score: 0.22222,
				},
			},
			Domain: &preyDomainScore{
				"prey1": &domainScoreComponents{
					conflictingDomains: []string{"domainX"},
					score:              0.45,
					supportingDomains:  []string{"domain1", "domain1", "domain2"},
				},
				"prey2": &domainScoreComponents{
					conflictingDomains: []string{"domainY", "domainZ"},
					score:              0.48,
					supportingDomains:  []string{"domain2", "domain3"},
				},
			},
			Study: &preyStudyScore{
				"prey1": &studyScoreComponents{
					fractionation: []string{"GO:111111", "GO:222222"},
					hpa:           []string{},
					score:         0.5,
				},
				"prey2": &studyScoreComponents{
					fractionation: []string{"GO:333333"},
					hpa:           []string{"GO:444444"},
					score:         1,
				},
			},
			Text: &preyTextScore{
				"prey1": &textScoreComponents{
					GOID:  "GO:111111",
					score: 0.75,
				},
				"prey2": &textScoreComponents{
					GOID:  "GO:444444",
					score: 0.25,
				},
			},
		}
		inputFiles := fileContent{
			predictions:       predictions,
			predictionSummary: predictionSummary,
		}

		expected := "prey\tcompartment\tGO term(s)\tGO ID(s)\tbait component\tstudy component\ttext component\tdomain component\ttotal score\tbaits\tHPA supporting\tFractionation supporting\tbest text term\tsupporting domains\tconflicting domains\n" +
			"prey1\t1\tTerm 1\tGO:111111\t0.11111\t0.50000\t0.75000\t0.45000\t0.45370\tbait1\t\tGO:111111;GO:222222\tGO:111111\tdomain1;domain1;domain2\tdomainX\n" +
			"prey2\t2\tTerm 2\tGO:222222\t0.22222\t1.00000\t0.25000\t0.48000\t0.49074\tbait2;bait3\tGO:444444\tGO:333333\tGO:444444\tdomain2;domain3\tdomainY;domainZ\n"

		writeScores(scores, inputFiles, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
