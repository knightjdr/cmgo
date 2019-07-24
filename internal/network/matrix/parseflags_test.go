package matrix

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
})

var _ = Describe("Parseflags", func() {
	Context("all command line arguments", func() {
		It("should parse arguments", func() {
			os.Args = []string{
				"cmd",
				"-colorList", "colors.txt",
				"-cutoff", "0.5",
				"-localizations", "localizations.txt",
				"-matrix", "node-profiles.txt",
				"-nodeLocalizations", "node-localizations.txt",
				"-outFile", "out.txt",
				"-outFileNetwork", "out.cyjs",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				colorList:         "colors.txt",
				cutoff:            0.5,
				localizations:     "localizations.txt",
				matrix:            "node-profiles.txt",
				nodeLocalizations: "node-localizations.txt",
				outFile:           "out.txt",
				outFileNetwork:    "out.cyjs",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})

	Context("only required command line arguments", func() {
		It("should set defaults", func() {
			os.Args = []string{
				"cmd",
				"-colorList", "colors.txt",
				"-localizations", "localizations.txt",
				"-matrix", "node-profiles.txt",
				"-nodeLocalizations", "node-localizations.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.cutoff).To(Equal(0.01), "should set default cutoff")
			Expect(options.outFile).To(Equal("corr.txt"), "should set default out file")
			Expect(options.outFileNetwork).To(Equal("corr.cyjs"), "should set default network out file")
		})
	})

	Context("missing required command line arguments", func() {
		It("should set defaults", func() {
			os.Args = []string{
				"cmd",
			}
			fileOptions := map[string]interface{}{}

			_, err := parseFlags(fileOptions)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("argument passed via input file", func() {
		It("should set defaults", func() {
			os.Args = []string{
				"cmd",
			}
			fileOptions := map[string]interface{}{
				"colorList":         "file-colors.txt",
				"cutoff":            0.5,
				"localizations":     "file-localizations.txt",
				"matrix":            "file-node-profiles.txt",
				"nodeLocalizations": "file-node-localizations.txt",
				"outFile":           "file-out.txt",
				"outFileNetwork":    "file-out.cyjs",
			}

			expected := parameters{
				colorList:         "file-colors.txt",
				cutoff:            0.5,
				localizations:     "file-localizations.txt",
				matrix:            "file-node-profiles.txt",
				nodeLocalizations: "file-node-localizations.txt",
				outFile:           "file-out.txt",
				outFileNetwork:    "file-out.cyjs",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
