package prediction

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
				"-baitExpected", "bait-expected.txt",
				"-fdr", "0.05",
				"-goHierarchy", "go.obo",
				"-outFile", "out.txt",
				"-predictions", "predictions.txt",
				"-predictionSummary", "prediction-summary.txt",
				"-predictionType", "safe",
				"-saint", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				baitExpected:      "bait-expected.txt",
				fdr:               0.05,
				goHierarchy:       "go.obo",
				outFile:           "out.txt",
				predictions:       "predictions.txt",
				predictionSummary: "prediction-summary.txt",
				predictionType:    "safe",
				saint:             "saint.txt",
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
				"-baitExpected", "bait-expected.txt",
				"-goHierarchy", "go.obo",
				"-predictions", "predictions.txt",
				"-predictionSummary", "prediction-summary.txt",
				"-saint", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.fdr).To(Equal(0.01), "should set default FDR")
			Expect(options.outFile).To(Equal("prediction-score.txt"), "should set output file name")
			Expect(options.predictionType).To(Equal("nmf"), "should set default prediction type")
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
				"baitExpected":      "file-bait-expected.txt",
				"fdr":               0.05,
				"goHierarchy":       "file-go.obo",
				"outFile":           "file-out.txt",
				"predictions":       "file-predictions.txt",
				"predictionSummary": "file-prediction-summary.txt",
				"predictionType":    "safe",
				"saint":             "file-saint.txt",
			}

			expected := parameters{
				baitExpected:      "file-bait-expected.txt",
				fdr:               0.05,
				goHierarchy:       "file-go.obo",
				outFile:           "file-out.txt",
				predictions:       "file-predictions.txt",
				predictionSummary: "file-prediction-summary.txt",
				predictionType:    "safe",
				saint:             "file-saint.txt",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
