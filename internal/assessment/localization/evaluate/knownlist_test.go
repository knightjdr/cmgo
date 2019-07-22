package evaluate

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write gene list with known status", func() {
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

		expected := "gene\tterm(s)\tID(s)\tknown\n" +
			"a\t1;2\tGO:1;GO:2\ttrue\n" +
			"b\t2\tGO:2\ttrue\n" +
			"c\t1\tGO:1\tfalse\n" +
			"d\t3\tGO:3\ttrue\n" +
			"e\t2;3\tGO:2;GO:3\ttrue\n" +
			"f\t1\tGO:1\ttrue\n" +
			"g\t\t\tfalse\n" +
			"h\t3\tGO:3\tfalse\n"
		knownList(genes, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
