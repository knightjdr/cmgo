package rankmetrics

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
				"-fasta", "database.fasta",
				"-fdr", "0.05",
				"-outFile", "out.txt",
				"-saint", "saint.txt",
				"-turnoverFile", "turnover-rates.txt",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				fasta:        "database.fasta",
				fdr:          0.05,
				outFile:      "out.txt",
				saint:        "saint.txt",
				turnoverFile: "turnover-rates.txt",
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
				"-fasta", "database.fasta",
				"-saint", "saint.txt",
				"-turnoverFile", "turnover-rates.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.fdr).To(Equal(0.01), "should set default FDR")
			Expect(options.outFile).To(Equal("turnover-by-rank.txt"), "should set default out file")
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
				"fasta":        "file-database.txt",
				"fdr":          "0.05",
				"outFile":      "file-out.txt",
				"saint":        "file-saint.txt",
				"turnoverFile": "file-turnover-rates.txt",
			}

			expected := parameters{
				fasta:        "file-database.txt",
				fdr:          0.05,
				outFile:      "file-out.txt",
				saint:        "file-saint.txt",
				turnoverFile: "file-turnover-rates.txt",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
