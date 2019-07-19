package evaluate

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write summary", func() {
	It("should write to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		genes := map[string]map[string]map[string]interface{}{
			"a": map[string]map[string]interface{}{
				"GO:1": map[string]interface{}{
					"known": false,
					"name":  "1",
				},
				"GO:2": map[string]interface{}{
					"known": true,
					"name":  "2",
				},
			},
			"b": map[string]map[string]interface{}{
				"GO:2": map[string]interface{}{
					"known": true,
					"name":  "2",
				},
			},
			"c": map[string]map[string]interface{}{
				"GO:1": map[string]interface{}{
					"known": false,
					"name":  "1",
				},
			},
			"d": map[string]map[string]interface{}{
				"GO:3": map[string]interface{}{
					"known": true,
					"name":  "3",
				},
			},
			"e": map[string]map[string]interface{}{
				"GO:3": map[string]interface{}{
					"known": true,
					"name":  "3",
				},
				"GO:2": map[string]interface{}{
					"known": true,
					"name":  "2",
				},
			},
			"f": map[string]map[string]interface{}{
				"GO:1": map[string]interface{}{
					"known": true,
					"name":  "1",
				},
			},
			"g": map[string]map[string]interface{}{
				"": map[string]interface{}{
					"known": false,
					"name":  "",
				},
			},
			"h": map[string]map[string]interface{}{
				"GO:3": map[string]interface{}{
					"known": false,
					"name":  "3",
				},
			},
		}

		expected := "total genes: 8\n" +
			"known genes: 5\n" +
			"fraction known: 0.6250\n\n" +
			"GO ID\tterm\ttotal\tknown\tfraction\n" +
			"-\tunknown\t1\t0\t0.0000\n" +
			"GO:1\t1\t3\t1\t0.3333\n" +
			"GO:2\t2\t3\t3\t1.0000\n" +
			"GO:3\t3\t3\t2\t0.6667\n"
		summarize(genes, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
