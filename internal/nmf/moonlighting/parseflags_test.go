package moonlighting

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
				"-dissimilarityFile", "dissimilarity.txt",
				"-minimumNmfScore", "0.05",
				"-nmfBasis", "nmf.txt",
				"-outFile", "out.txt",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				dissimilarityFile: "dissimilarity.txt",
				minimumNmfScore:   0.05,
				nmfBasis:          "nmf.txt",
				outFile:           "out.txt",
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
				"-dissimilarityFile", "dissimilarity.txt",
				"-nmfBasis", "nmf.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.minimumNmfScore).To(Equal(0.15), "should set default minimum NMF score")
			Expect(options.outFile).To(Equal("moonlighting.txt"), "should set default out file")
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
				"dissimilarityFile": "file-dissimilarity.txt",
				"minimumNmfScore":   "0.05",
				"nmfBasis":          "file-nmf.txt",
				"outFile":           "file-out.txt",
			}

			expected := parameters{
				dissimilarityFile: "file-dissimilarity.txt",
				minimumNmfScore:   0.05,
				nmfBasis:          "file-nmf.txt",
				outFile:           "file-out.txt",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
