package localize

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
				"database", "database.fasta",
				"-fdr", "0.02",
				"-goAnnotations", "annotations.gaf",
				"-goHierarchy", "hierarchy.obo",
				"-minBaits", "2",
				"-namespace", "BP",
				"-outFile", "out.txt",
				"-saintFile", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				database:      "database.fasta",
				fdr:           0.02,
				goAnnotations: "annotations.gaf",
				goHierarchy:   "hierarchy.obo",
				minBaits:      2,
				namespace:     "BP",
				outFile:       "out.txt",
				saintFile:     "saint.txt",
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
				"-database", "database.fasta",
				"-goAnnotations", "annotations.gaf",
				"-goHierarchy", "hierarchy.obo",
				"-saintFile", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.fdr).To(Equal(0.01), "should set default FDR")
			Expect(options.minBaits).To(Equal(1), "should set default minimum baits")
			Expect(options.namespace).To(Equal("CC"), "should set default GO namespace")
			Expect(options.outFile).To(Equal("lba-localization.txt"), "should set default out file")
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
				"database":      "file-database.fasta",
				"fdr":           0.02,
				"goAnnotations": "file-annotations.gaf",
				"goHierarchy":   "file-hierarchy.obo",
				"minBaits":      2,
				"namespace":     "BP",
				"outFile":       "file-out.txt",
				"saintFile":     "file-saint.txt",
			}

			expected := parameters{
				database:      "file-database.fasta",
				fdr:           0.02,
				goAnnotations: "file-annotations.gaf",
				goHierarchy:   "file-hierarchy.obo",
				minBaits:      2,
				namespace:     "BP",
				outFile:       "file-out.txt",
				saintFile:     "file-saint.txt",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
