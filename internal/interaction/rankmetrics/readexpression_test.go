package rankmetrics

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var gixText = `[
	{
		"uniprot": ["Q13685", "Q00001"],
		"protein-expression": {
			"cells": {
				"HEK-293": {
					"intensity": 1.2
				}
			}
		}
	},
	{
		"uniprot": ["P123A4"],
		"protein-expression": {
			"cells": {
				"HEK-293": {
					"intensity": 2.4
				}
			}
		}
	},
	{
		"uniprot": ["P42765", "Q00003"],
		"protein-expression": {
			"cells": {
				"HEK-293": {
					"intensity": 3.14
				}
			}
		}
	}
]`

var _ = Describe("Read gix database", func() {
	It("should read gene names with protein expression data", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/gix.json",
			[]byte(gixText),
			0444,
		)

		data := analysis{
			parameters: parameters{
				gixDB: "test/gix.json",
			},
			uniprotMapping: map[string]string{
				"Q13685": "AAMP",
				"P123A4": "AARS2",
				"P42765": "ACAA2",
				"Q00003": "ACAA2",
			},
		}

		expected := map[string]float64{
			"AAMP":  1.2,
			"AARS2": 2.4,
			"ACAA2": 3.14,
		}
		Expect(readExpression(data)).To(Equal(expected))
	})
})
