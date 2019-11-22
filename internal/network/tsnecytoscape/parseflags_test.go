package tsnecytoscape

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parseflags", func() {
	Context("all command line arguments", func() {
		It("should parse arguments", func() {
			os.Args = []string{
				"cmd",
				"-colorList", "colors.txt",
				"-localizations", "localizations.txt",
				"-nodeCoordinates", "coordinates.txt",
				"-nodeLocalizations", "node-localizations.txt",
				"-outFile", "out.svg",
				"-width", "500",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				colorList:         "colors.txt",
				localizations:     "localizations.txt",
				nodeCoordinates:   "coordinates.txt",
				nodeLocalizations: "node-localizations.txt",
				outFile:           "out.svg",
				width:             500,
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
				"-nodeCoordinates", "coordinates.txt",
				"-nodeLocalizations", "node-localizations.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.outFile).To(Equal("map.cyjs"), "should set default out file")
			Expect(options.width).To(Equal(float64(1000)), "should set default width")
		})
	})

	Context("missing required command line arguments", func() {
		It("should report error", func() {
			os.Args = []string{
				"cmd",
			}
			fileOptions := map[string]interface{}{}

			_, err := parseFlags(fileOptions)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("argument passed via input file", func() {
		It("should set variables from file", func() {
			os.Args = []string{
				"cmd",
			}
			fileOptions := map[string]interface{}{
				"colorList":         "file-colors.txt",
				"localizations":     "file-localizations.txt",
				"nodeCoordinates":   "file-coordinates.txt",
				"nodeLocalizations": "file-node-localizations.txt",
				"outFile":           "file-out.svg",
				"width":             "500",
			}
			expected := parameters{
				colorList:         "file-colors.txt",
				localizations:     "file-localizations.txt",
				nodeCoordinates:   "file-coordinates.txt",
				nodeLocalizations: "file-node-localizations.txt",
				outFile:           "file-out.svg",
				width:             500,
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
