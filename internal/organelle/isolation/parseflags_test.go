package isolation

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
				"-abundanceCap", "500",
				"-basisMatrix", "basis.txt",
				"-correlationCutoff", "0.75",
				"-nmfLocalization", "localization.txt",
				"-nmfSummary", "summary.txt",
				"-outFile", "out.txt",
				"-svgFile", "out.svg",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				abundanceCap:      500,
				basisMatrix:       "basis.txt",
				correlationCutoff: 0.75,
				nmfLocalization:   "localization.txt",
				nmfSummary:        "summary.txt",
				outFile:           "out.txt",
				svgFile:           "out.svg",
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
				"-basisMatrix", "basis.txt",
				"-nmfLocalization", "localization.txt",
				"-nmfSummary", "summary.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.abundanceCap).To(Equal(float64(1000)), "should set default abundance cap")
			Expect(options.correlationCutoff).To(Equal(0.9), "should set default correlation cutoff")
			Expect(options.outFile).To(Equal("organelle-isolation.txt"), "should set default out file")
			Expect(options.svgFile).To(Equal("organelle-isolation.svg"), "should set default svg file")
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
				"abundanceCap":      "500",
				"basisMatrix":       "file-basis.txt",
				"correlationCutoff": "0.75",
				"nmfLocalization":   "file-localization.txt",
				"nmfSummary":        "file-summary.txt",
				"outFile":           "file-out.txt",
				"svgFile":           "file-out.svg",
			}

			expected := parameters{
				abundanceCap:      500,
				basisMatrix:       "file-basis.txt",
				correlationCutoff: 0.75,
				nmfLocalization:   "file-localization.txt",
				nmfSummary:        "file-summary.txt",
				outFile:           "file-out.txt",
				svgFile:           "file-out.svg",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
