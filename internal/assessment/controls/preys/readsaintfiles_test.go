package preys

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var baitText = `128_7909	7909_BirAFLAG_April2017	C	bira-flag
128_8301	8301_FLAG_alone	C	empty`

var interText = `
128_7909	7909_BirAFLAG_April2017	BirA_R118G_H0QFJ5	1075	25
128_7909	7909_BirAFLAG_April2017	NP_000029.2	5	4
128_8301	8301_FLAG_alone	NP_000048.1	4	4
128_8301	8301_FLAG_alone	NP_000108.1	7	6`

var preyText = `BirA_R118G_H0QFJ5	321	BirA_R118G_H0QFJ5
NP_000029.2	2843	APC
NP_000048.1	1417	BLM
NP_000108.1	254	EMD`

var _ = Describe("Read SAINT input files", func() {
	It("should parse baits with interactions", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(fs.Instance, "test/bait.dat", []byte(baitText), 0444)
		afero.WriteFile(fs.Instance, "test/inter.dat", []byte(interText), 0444)
		afero.WriteFile(fs.Instance, "test/prey.dat", []byte(preyText), 0444)

		options := parameters{
			bait:  "test/bait.dat",
			inter: "test/inter.dat",
			prey:  "test/prey.dat",
		}

		expectedBaits := map[string]string{
			"128_7909": "bira-flag",
			"128_8301": "empty",
		}
		expectedInteractions := map[string]map[string]int{
			"128_7909": map[string]int{
				"BirA_R118G_H0QFJ5": 1075,
				"APC":               5,
			},
			"128_8301": map[string]int{
				"BLM": 4,
				"EMD": 7,
			},
		}

		actualBaits, actualInteractions := readSaintFiles(options)
		Expect(actualBaits).To(Equal(expectedBaits), "should parse bait IDs with control type from file")
		Expect(actualInteractions).To(Equal(expectedInteractions), "should parse control interactions")
	})
})
