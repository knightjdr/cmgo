package correlation

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write cyjs", func() {
	It("should write json to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		colors := []string{"#000", "#f00", "#0f0", "#00f"}
		corr := [][]float64{
			{1, -0.2, 0.71, 0.82},
			{-0.2, 1, 0.95, 0.3},
			{0.71, 0.95, 1, 0.71},
			{0.82, 0.3, 0.71, 1},
		}
		genes := []string{"a", "b", "c", "d"}
		nodeLocalizations := map[string]map[string]string{
			"a": map[string]string{
				"GO:1": "1",
				"GO:2": "2",
			},
			"b": map[string]string{
				"GO:3": "3",
			},
			"c": map[string]string{
				"GO:2": "2",
			},
			"d": map[string]string{},
		}
		possibleLocalizations := []string{"GO:2", "GO:1", "GO:3"}

		expected := "{\n" +
			"\t\"elements\": {\n" +
			"\t\t\"nodes\": [\n" +
			"\t\t\t{\n" +
			"\t\t\t\t\"data\": {\n" +
			"\t\t\t\t\t\"color\": \"#0f0\",\n" +
			"\t\t\t\t\t\"compartment\": \"1\",\n" +
			"\t\t\t\t\t\"id\": \"a\",\n" +
			"\t\t\t\t\t\"name\": \"a\"\n" +
			"\t\t\t\t},\n" +
			"\t\t\t\t\"position\": {\n" +
			"\t\t\t\t\t\"x\": 1,\n" +
			"\t\t\t\t\t\"y\": 1\n" +
			"\t\t\t\t}\n" +
			"\t\t\t},\n" +
			"\t\t\t{\n" +
			"\t\t\t\t\"data\": {\n" +
			"\t\t\t\t\t\"color\": \"#00f\",\n" +
			"\t\t\t\t\t\"compartment\": \"3\",\n" +
			"\t\t\t\t\t\"id\": \"b\",\n" +
			"\t\t\t\t\t\"name\": \"b\"\n" +
			"\t\t\t\t},\n" +
			"\t\t\t\t\"position\": {\n" +
			"\t\t\t\t\t\"x\": 1,\n" +
			"\t\t\t\t\t\"y\": 1\n" +
			"\t\t\t\t}\n" +
			"\t\t\t},\n" +
			"\t\t\t{\n" +
			"\t\t\t\t\"data\": {\n" +
			"\t\t\t\t\t\"color\": \"#f00\",\n" +
			"\t\t\t\t\t\"compartment\": \"2\",\n" +
			"\t\t\t\t\t\"id\": \"c\",\n" +
			"\t\t\t\t\t\"name\": \"c\"\n" +
			"\t\t\t\t},\n" +
			"\t\t\t\t\"position\": {\n" +
			"\t\t\t\t\t\"x\": 1,\n" +
			"\t\t\t\t\t\"y\": 1\n" +
			"\t\t\t\t}\n" +
			"\t\t\t},\n" +
			"\t\t\t{\n" +
			"\t\t\t\t\"data\": {\n" +
			"\t\t\t\t\t\"color\": \"#000\",\n" +
			"\t\t\t\t\t\"compartment\": \"\",\n" +
			"\t\t\t\t\t\"id\": \"d\",\n" +
			"\t\t\t\t\t\"name\": \"d\"\n" +
			"\t\t\t\t},\n" +
			"\t\t\t\t\"position\": {\n" +
			"\t\t\t\t\t\"x\": 1,\n" +
			"\t\t\t\t\t\"y\": 1\n" +
			"\t\t\t\t}\n" +
			"\t\t\t}\n" +
			"\t\t],\n" +
			"\t\t\"edges\": [\n" +
			"\t\t\t{\n" +
			"\t\t\t\t\"data\": {\n" +
			"\t\t\t\t\t\"distance\": 0.71,\n" +
			"\t\t\t\t\t\"id\": \"a-c\",\n" +
			"\t\t\t\t\t\"name\": \"a (interacts with) c\",\n" +
			"\t\t\t\t\t\"source\": \"a\",\n" +
			"\t\t\t\t\t\"target\": \"c\"\n" +
			"\t\t\t\t}\n" +
			"\t\t\t},\n" +
			"\t\t\t{\n" +
			"\t\t\t\t\"data\": {\n" +
			"\t\t\t\t\t\"distance\": 0.82,\n" +
			"\t\t\t\t\t\"id\": \"a-d\",\n" +
			"\t\t\t\t\t\"name\": \"a (interacts with) d\",\n" +
			"\t\t\t\t\t\"source\": \"a\",\n" +
			"\t\t\t\t\t\"target\": \"d\"\n" +
			"\t\t\t\t}\n" +
			"\t\t\t},\n" +
			"\t\t\t{\n" +
			"\t\t\t\t\"data\": {\n" +
			"\t\t\t\t\t\"distance\": 0.95,\n" +
			"\t\t\t\t\t\"id\": \"b-c\",\n" +
			"\t\t\t\t\t\"name\": \"b (interacts with) c\",\n" +
			"\t\t\t\t\t\"source\": \"b\",\n" +
			"\t\t\t\t\t\"target\": \"c\"\n" +
			"\t\t\t\t}\n" +
			"\t\t\t},\n" +
			"\t\t\t{\n" +
			"\t\t\t\t\"data\": {\n" +
			"\t\t\t\t\t\"distance\": 0.71,\n" +
			"\t\t\t\t\t\"id\": \"c-d\",\n" +
			"\t\t\t\t\t\"name\": \"c (interacts with) d\",\n" +
			"\t\t\t\t\t\"source\": \"c\",\n" +
			"\t\t\t\t\t\"target\": \"d\"\n" +
			"\t\t\t\t}\n" +
			"\t\t\t}\n" +
			"\t\t]\n" +
			"\t}\n" +
			"}"

		writeJSON(corr, genes, 0.71, nodeLocalizations, possibleLocalizations, colors, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
