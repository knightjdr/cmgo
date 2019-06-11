package gradient

import (
	"testing"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestWriteGradient(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)

	similarity := []int{0, 1, 2, 2, 1, 1}
	wanted := "<svg xmlns=\"http://www.w3.org/2000/svg\" height=\"14\" width=\"22\" viewBox=\"0 0 22 14\">\n" +
		"\t<rect x=\"1\" y=\"1\" height=\"2\" width=\"20\" fill=\"#fff\" />\n" +
		"\t<rect x=\"1\" y=\"3\" height=\"2\" width=\"20\" fill=\"#ff8c66\" />\n" +
		"\t<rect x=\"1\" y=\"5\" height=\"2\" width=\"20\" fill=\"#cc3300\" />\n" +
		"\t<rect x=\"1\" y=\"7\" height=\"2\" width=\"20\" fill=\"#cc3300\" />\n" +
		"\t<rect x=\"1\" y=\"9\" height=\"2\" width=\"20\" fill=\"#ff8c66\" />\n" +
		"\t<rect x=\"1\" y=\"11\" height=\"2\" width=\"20\" fill=\"#ff8c66\" />\n" +
		"\t<rect x=\"0\" y=\"0\" height=\"14\" width=\"22\" fill=\"none\" stroke=\"black\" stroke-width=\"1\" />\n" +
		"</svg>\n"
	writeGradient(similarity, "test/out.svg")
	bytes, _ := afero.ReadFile(fs.Instance, "test/out.svg")
	assert.Equal(t, wanted, string(bytes), "Should write gradient to file")
}
