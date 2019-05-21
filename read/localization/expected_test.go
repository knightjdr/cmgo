package localization

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var expectedText = `id	bait	localization
1	AARS2	mitochondrial matrix
2	ACBD5	peroxisome
3	ACTB	actin cytoskeleton
12	ANAPC2	"cytoplasm;nucleoplasm;nucleus"
13	ANK3	"cell junction;plasma membrane"
`

func TestMapExpectedLine(t *testing.T) {
	// TEST1: single localization
	line := []string{"5", "baitA", "membrane"}
	wantedBait := "baitA"
	wantedInfo := ExpectedLocalization{
		ID:    5,
		Terms: []string{"membrane"},
	}
	resultBait, resultInfo := mapExpectedLine(line)
	assert.Equal(t, wantedBait, resultBait, "Should read bait name from line")
	assert.Equal(t, wantedInfo, resultInfo, "Should read bait localization information from line")

	// TEST2: multiple localizations
	line = []string{"6", "baitB", "membrane;nucleus"}
	wantedBait = "baitB"
	wantedInfo = ExpectedLocalization{
		ID:    6,
		Terms: []string{"membrane", "nucleus"},
	}
	resultBait, resultInfo = mapExpectedLine(line)
	assert.Equal(t, wantedBait, resultBait, "Should read bait name from line")
	assert.Equal(t, wantedInfo, resultInfo, "Should read multiple bait localizations from line")

	// TEST3: multiple localizations with leading and trailing quotes
	line = []string{"7", "baitC", "\"membrane;nucleus\""}
	wantedBait = "baitC"
	wantedInfo = ExpectedLocalization{
		ID:    7,
		Terms: []string{"membrane", "nucleus"},
	}
	resultBait, resultInfo = mapExpectedLine(line)
	assert.Equal(t, wantedBait, resultBait, "Should read bait name from line")
	assert.Equal(t, wantedInfo, resultInfo, "Should read multiple bait localizations from line")
}

func TestExpected(t *testing.T) {
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
	wanted := ExpectedLocalizations{
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
	assert.Equal(t, wanted, Expected("test/expected.txt"), "Should read a file with expected localizations")
}
