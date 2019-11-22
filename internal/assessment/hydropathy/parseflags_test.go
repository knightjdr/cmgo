package hydropathy

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
				"-bioplexFile", "bioplex.txt",
				"database", "database.fasta",
				"-fdr", "0.02",
				"-saintFile", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				bioplexFile: "bioplex.txt",
				database:    "database.fasta",
				fdr:         0.02,
				saintFile:   "saint.txt",
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
				"-bioplexFile", "bioplex.txt",
				"-database", "database.fasta",
				"-saintFile", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.fdr).To(Equal(0.01), "should set default FDR")
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
				"bioplexFile": "file-bioplex.txt",
				"database":    "file-database.fasta",
				"fdr":         0.02,
				"saintFile":   "file-saint.txt",
			}

			expected := parameters{
				bioplexFile: "file-bioplex.txt",
				database:    "file-database.fasta",
				fdr:         0.02,
				saintFile:   "file-saint.txt",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
