package localization

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var expectedText = `id	bait	localization
1	AARS2	mitochondrial matrix
2	ACBD5	peroxisome
3	ACTB	actin cytoskeleton
12	ANAPC2	"cytoplasm;nucleoplasm;nucleus"
13	ANK3	"cell junction;plasma membrane"
`

var _ = Describe("Map expected bait localization line", func() {
	It("should read line with a single localization", func() {
		line := []string{"5", "baitA", "membrane"}
		expected := ExpectedLocalization{
			ID:    5,
			Terms: []string{"membrane"},
		}
		resultBait, resultInfo := mapExpectedLine(line)
		Expect(resultBait).To(Equal("baitA"), "Should read bait name from line")
		Expect(resultInfo).To(Equal(expected), "Should read bait localization information from line")
	})

	It("should read line with multiple localizations", func() {
		line := []string{"6", "baitB", "membrane;nucleus"}
		expected := ExpectedLocalization{
			ID:    6,
			Terms: []string{"membrane", "nucleus"},
		}
		resultBait, resultInfo := mapExpectedLine(line)
		Expect(resultBait).To(Equal("baitB"), "Should read bait name from line")
		Expect(resultInfo).To(Equal(expected), "Should read multiple bait localizations from line")
	})

	It("should read line with multiple localizations with leading and trailing quotes", func() {
		line := []string{"7", "baitC", "\"membrane;nucleus\""}
		expected := ExpectedLocalization{
			ID:    7,
			Terms: []string{"membrane", "nucleus"},
		}
		resultBait, resultInfo := mapExpectedLine(line)
		Expect(resultBait).To(Equal("baitC"), "Should read bait name from line")
		Expect(resultInfo).To(Equal(expected), "Should read multiple bait localizations from line")
	})
})

var _ = Describe("Read expected bait localizations", func() {
	It("should read file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/expected.txt",
			[]byte(expectedText),
			0444,
		)

		// TEST
		expected := ExpectedLocalizations{
			"AARS2": ExpectedLocalization{
				ID:    1,
				Terms: []string{"mitochondrial matrix"},
			},
			"ACBD5": ExpectedLocalization{
				ID:    2,
				Terms: []string{"peroxisome"},
			},
			"ACTB": ExpectedLocalization{
				ID:    3,
				Terms: []string{"actin cytoskeleton"},
			},
			"ANAPC2": ExpectedLocalization{
				ID:    12,
				Terms: []string{"cytoplasm", "nucleoplasm", "nucleus"},
			},
			"ANK3": ExpectedLocalization{
				ID:    13,
				Terms: []string{"cell junction", "plasma membrane"},
			},
		}
		Expect(Expected("test/expected.txt")).To(Equal(expected))
	})
})
