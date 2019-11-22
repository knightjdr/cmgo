package uv

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
				"-basisMatrix", "basis.csv",
				"-goAnnotations", "annotations.gaf",
				"-goHierarchy", "hierarchy.obo",
				"-maxGenesPerRank", "50",
				"-minRankValue", "0.5",
				"-namespace", "BP",
				"-nmfLocalization", "nmf-localizations.txt",
				"-nmfSummary", "rank-summary.txt",
				"-outFile", "out.txt",
				"-withinRankMax", "0.6",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				basisMatrix:     "basis.csv",
				goAnnotations:   "annotations.gaf",
				goHierarchy:     "hierarchy.obo",
				maxGenesPerRank: 50,
				minRankValue:    0.5,
				namespace:       "BP",
				nmfLocalization: "nmf-localizations.txt",
				nmfSummary:      "rank-summary.txt",
				outFile:         "out.txt",
				withinRankMax:   0.6,
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
				"-basisMatrix", "basis.csv",
				"-goAnnotations", "annotations.gaf",
				"-goHierarchy", "hierarchy.obo",
				"-nmfLocalization", "nmf-localizations.txt",
				"-nmfSummary", "rank-summary.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.maxGenesPerRank).To(Equal(100), "should set default maxGenesPerRank")
			Expect(options.outFile).To(Equal("uv-assessment.txt"), "should set default out file")
			Expect(options.withinRankMax).To(Equal(0.75), "should set default withinRankMax")
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
				"basisMatrix":     "file-basis.csv",
				"goAnnotations":   "file-annotations.gaf",
				"goHierarchy":     "file-hierarchy.obo",
				"namespace":       "BP",
				"maxGenesPerRank": 50,
				"minRankValue":    0.5,
				"nmfLocalization": "file-nmf-localizations.txt",
				"nmfSummary":      "file-rank-summary.txt",
				"outFile":         "file-out.txt",
				"withinRankMax":   0.6,
			}

			expected := parameters{
				basisMatrix:     "file-basis.csv",
				goAnnotations:   "file-annotations.gaf",
				goHierarchy:     "file-hierarchy.obo",
				maxGenesPerRank: 50,
				minRankValue:    0.5,
				namespace:       "BP",
				nmfLocalization: "file-nmf-localizations.txt",
				nmfSummary:      "file-rank-summary.txt",
				outFile:         "file-out.txt",
				withinRankMax:   0.6,
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
