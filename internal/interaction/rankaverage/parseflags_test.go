package rankaverage

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
				"-fdr", "0.05",
				"-outFile", "out.txt",
				"-preys", "preys.txt",
				"-saint", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				fdr:     0.05,
				outFile: "out.txt",
				preys:   "preys.txt",
				saint:   "saint.txt",
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
				"-preys", "preys.txt",
				"-saint", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.fdr).To(Equal(0.01), "should set default FDR")
			Expect(options.outFile).To(Equal("prey-rank-average.txt"), "should set default out file")
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
				"fdr":     "0.05",
				"outFile": "file-out.txt",
				"preys":   "file-preys.txt",
				"saint":   "file-saint.txt",
			}

			expected := parameters{
				fdr:     0.05,
				outFile: "file-out.txt",
				preys:   "file-preys.txt",
				saint:   "file-saint.txt",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
