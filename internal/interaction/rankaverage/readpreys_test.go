package rankaverage

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var preyText = `gene
preyA
preyB
preyC
`

var _ = Describe("Read prey list", func() {
	It("should read gene names", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/preys.txt",
			[]byte(preyText),
			0444,
		)

		expected := []string{"preyA", "preyB", "preyC"}
		Expect(readPreys("test/preys.txt")).To(Equal(expected))
	})
})
