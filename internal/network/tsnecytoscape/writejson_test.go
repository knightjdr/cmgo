package tsnecytoscape

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/tsne"
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
		nodeCoordinates := map[string]tsne.Coordinate{
			"a": tsne.Coordinate{X: -20, Y: 50},
			"b": tsne.Coordinate{X: 0, Y: 100},
			"c": tsne.Coordinate{X: 20, Y: -30},
			"d": tsne.Coordinate{X: 0, Y: 0},
		}
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
		transformation := map[string]float64{
			"scale":      2,
			"translateX": 20,
			"translateY": 30,
		}

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
			"\t\t\t\t\t\"x\": 0,\n" +
			"\t\t\t\t\t\"y\": 160\n" +
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
			"\t\t\t\t\t\"x\": 40,\n" +
			"\t\t\t\t\t\"y\": 260\n" +
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
			"\t\t\t\t\t\"x\": 80,\n" +
			"\t\t\t\t\t\"y\": 0\n" +
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
			"\t\t\t\t\t\"x\": 40,\n" +
			"\t\t\t\t\t\"y\": 60\n" +
			"\t\t\t\t}\n" +
			"\t\t\t}\n" +
			"\t\t],\n" +
			"\t\t\"edges\": []\n" +
			"\t}\n" +
			"}"

		writeJSON(nodeCoordinates, nodeLocalizations, possibleLocalizations, colors, transformation, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
