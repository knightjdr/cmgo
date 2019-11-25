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
				"-basisMatrix", "basis.txt",
				"-dissimilarityMatrix", "dissimilarity.txt",
				"-minRankValue", "0.05",
				"-nmfSummary", "nmf-summary.txt",
				"-outFileHeatmap", "out-heatmap.svg",
				"-outFileMatrix", "out-matrix.txt",
				"-outFileScores", "out-scores.txt",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				basisMatrix:         "basis.txt",
				dissimilarityMatrix: "dissimilarity.txt",
				minRankValue:        0.05,
				nmfSummary:          "nmf-summary.txt",
				outFileHeatmap:      "out-heatmap.svg",
				outFileMatrix:       "out-matrix.txt",
				outFileScores:       "out-scores.txt",
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
				"-dissimilarityMatrix", "dissimilarity.txt",
				"-nmfSummary", "nmf-summary.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.minRankValue).To(Equal(0.15), "should set default minimum NMF score")
			Expect(options.outFileHeatmap).To(Equal("heatmap.svg"), "should set default out file for heatmap")
			Expect(options.outFileMatrix).To(Equal("matrix.txt"), "should set default out file for matrix")
			Expect(options.outFileScores).To(Equal("moonlighting.txt"), "should set default out file for moonlighting scores")
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
				"basisMatrix":         "file-basis.txt",
				"dissimilarityMatrix": "file-dissimilarity.txt",
				"minRankValue":        "0.05",
				"nmfSummary":          "file-nmf-summary.txt",
				"outFileHeatmap":      "file-heatmap.txt",
				"outFileMatrix":       "file-matrix.txt",
				"outFileScores":       "file-moonlighting.txt",
			}

			expected := parameters{
				basisMatrix:         "file-basis.txt",
				dissimilarityMatrix: "file-dissimilarity.txt",
				minRankValue:        0.05,
				nmfSummary:          "file-nmf-summary.txt",
				outFileHeatmap:      "file-heatmap.txt",
				outFileMatrix:       "file-matrix.txt",
				outFileScores:       "file-moonlighting.txt",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
