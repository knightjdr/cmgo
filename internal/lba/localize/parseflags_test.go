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
				"-enrichment", "enrichment.txt",
				"-localization", "localization.txt",
				"-outFilePrimary", "out.txt",
				"-outFileProfile", "profile.txt",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				enrichment:     "enrichment.txt",
				localization:   "localization.txt",
				outFilePrimary: "out.txt",
				outFileProfile: "profile.txt",
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
				"-enrichment", "enrichment.txt",
				"-localization", "localization.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.outFilePrimary).To(Equal("lba-primary.txt"), "should set default primary localization out file")
			Expect(options.outFileProfile).To(Equal("lba-profile.txt"), "should set default profile out file")
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
				"enrichment":     "file-enrichment.txt",
				"localization":   "file-localization.txt",
				"outFilePrimary": "file-out.txt",
				"outFileProfile": "file-profile.txt",
			}

			expected := parameters{
				enrichment:     "file-enrichment.txt",
				localization:   "file-localization.txt",
				outFilePrimary: "file-out.txt",
				outFileProfile: "file-profile.txt",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
