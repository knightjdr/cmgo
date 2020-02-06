package preys

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
				"-bait", "bait.dat",
				"-enrichmentLimit", "100",
				"-inter", "inter.dat",
				"-outFile", "out.txt",
				"-outFileEnrichment", "out-enrichment.txt",
				"-prey", "prey.dat",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				bait:              "bait.dat",
				enrichmentLimit:   100,
				inter:             "inter.dat",
				outFile:           "out.txt",
				outFileEnrichment: "out-enrichment.txt",
				prey:              "prey.dat",
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
				"-bait", "bait.dat",
				"-inter", "inter.dat",
				"-prey", "prey.dat",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.enrichmentLimit).To(Equal(200), "should set default enrichment limit")
			Expect(options.outFile).To(Equal("control-preys.txt"), "should set default out file")
			Expect(options.outFileEnrichment).To(Equal("enrichment.txt"), "should set default out file for enrichment")
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
			os.Args = []string{}
			fileOptions := map[string]interface{}{
				"bait":              "file-bait.dat",
				"enrichmentLimit":   100,
				"inter":             "file-inter.dat",
				"outFile":           "file-out.txt",
				"outFileEnrichment": "file-enrichment.txt",
				"prey":              "file-prey.dat",
			}

			expected := parameters{
				bait:              "file-bait.dat",
				enrichmentLimit:   100,
				inter:             "file-inter.dat",
				outFile:           "file-out.txt",
				outFileEnrichment: "file-enrichment.txt",
				prey:              "file-prey.dat",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
