package knownbyrank

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
				"-biogrid", "biogrid.txt",
				"-fdr", "0.05",
				"-intact", "intact.txt",
				"-outFile", "out.txt",
				"-saint", "saint.txt",
				"-species", "10000",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				biogrid: "biogrid.txt",
				fdr:     0.05,
				intact:  "intact.txt",
				outFile: "out.txt",
				saint:   "saint.txt",
				species: "10000",
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
				"-biogrid", "biogrid.txt",
				"-intact", "intact.txt",
				"-saint", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.fdr).To(Equal(0.01), "should set default FDR")
			Expect(options.outFile).To(Equal("known-by-rank.txt"), "should set default out file")
			Expect(options.species).To(Equal("9606"), "should set default species taxon ID")
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
				"biogrid": "file-biogrid.txt",
				"fdr":     "0.05",
				"intact":  "file-intact.txt",
				"outFile": "file-out.txt",
				"saint":   "file-saint.txt",
				"species": "10000",
			}

			expected := parameters{
				biogrid: "file-biogrid.txt",
				fdr:     0.05,
				intact:  "file-intact.txt",
				outFile: "file-out.txt",
				saint:   "file-saint.txt",
				species: "10000",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
