package prediction

import (
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
					score:              0.75,
					supportingDomains:  []string{"domain1", "domain1", "domain2"},
				},
				"prey2": &domainScoreComponents{
					conflictingDomains: []string{"domainY", "domainZ"},
					score:              0.25,
					supportingDomains:  []string{"domain2", "domain3"},
				},
			},
		}

		expected := "prey\tbait component\tdomain component\ttotal score\tbaits\tsupporting domains\tconflicting domains\n" +
			"prey1\t0.11111\t0.75000\t0.43056\tbait1\tdomain1;domain1;domain2\tdomainX\n" +
			"prey2\t0.22222\t0.25000\t0.23611\tbait2;bait3\tdomain2;domain3\tdomainY;domainZ\n"

		writeScores(scores, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
