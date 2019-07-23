package svg

import (
	"testing"

	"github.com/knightjdr/cmgo/internal/pkg/read/nmf"
	"github.com/knightjdr/cmgo/internal/pkg/read/tsne"
	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestWriteSVG(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)

	colors := []string{"#f00", "#000", "#ccc", "#fff"}
	coordinates := map[string]tsne.Coordinate{
		"a": tsne.Coordinate{X: 10, Y: 25},
		"b": tsne.Coordinate{X: 20, Y: 35},
		"c": tsne.Coordinate{X: 56, Y: 7},
	}
	localization := nmf.NMFlocalization{
		"a": nmf.GeneLocalization{Compartment: 3},
		"b": nmf.GeneLocalization{Compartment: 1},
		"c": nmf.GeneLocalization{Compartment: 2},
	}

	wanted := "<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"1000.00\" height=\"500.00\" viewbox=\"0 0 1000.00 500.00\">" +
		"<g id=\"network__zoom\" transform=\"translate(0, 0) scale(1)\">" +
		"<circle cx=\"10\" cy=\"25\" r=\"4px\" fill=\"#fff\" data-symbol=\"a\" />" +
		"<circle cx=\"20\" cy=\"35\" r=\"4px\" fill=\"#000\" data-symbol=\"b\" />" +
		"<circle cx=\"56\" cy=\"7\" r=\"4px\" fill=\"#ccc\" data-symbol=\"c\" />" +
		"</g>" +
		"</svg>"
	writeSVG(coordinates, colors, localization, 1000, 500, "test/out.txt")
	bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
	assert.Equal(t, wanted, string(bytes), "Should write network to svg file")
}
