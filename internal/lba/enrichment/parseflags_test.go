package enrichment

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
				"-database", "database.fasta",
				"-fdr", "0.02",
				"-minBaits", "2",
				"-minFC", "1",
				"-namespace", "BP",
				"-outFile", "out.txt",
				"-saintFile", "saint.txt",
				"-preyLimit", "50",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				database:  "database.fasta",
				fdr:       0.02,
				minBaits:  2,
				minFC:     1,
				namespace: "BP",
				outFile:   "out.txt",
				saintFile: "saint.txt",
				preyLimit: 50,
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
			Expect(options.minFC).To(Equal(float64(1)), "should set default minimum fold change")
			Expect(options.namespace).To(Equal("CC"), "should set default GO namespace")
			Expect(options.outFile).To(Equal("lba-enrichment.txt"), "should set default out file")
			Expect(options.preyLimit).To(Equal(100), "should set default number of preys to use for enrichment")
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
				"database":  "file-database.fasta",
				"fdr":       0.02,
				"minBaits":  2,
				"minFC":     1,
				"namespace": "BP",
				"outFile":   "file-out.txt",
				"saintFile": "file-saint.txt",
				"preyLimit": 50,
			}

			expected := parameters{
				database:  "file-database.fasta",
				fdr:       0.02,
				minBaits:  2,
				minFC:     1,
				namespace: "BP",
				outFile:   "file-out.txt",
				saintFile: "file-saint.txt",
				preyLimit: 50,
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
