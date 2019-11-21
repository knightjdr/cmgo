package interactions

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Read interaction file", func() {
	It("should read interactors", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/biogrid.txt",
			[]byte(biogridText),
			0444,
		)
		afero.WriteFile(
			fs.Instance,
			"test/intact.txt",
			[]byte(intactText),
			0444,
		)

		expected := map[string][]string{
			"ACVR1":  []string{"FNTA"},
			"ALDOA":  []string{"XRN1"},
			"AMPH":   []string{"SYNJ1"},
			"ARF3":   []string{"ARFIP1", "ARFIP2"},
			"ARFIP1": []string{"ARF3"},
			"ARFIP2": []string{"ARF3"},
			"FLNC":   []string{"MAP2K4"},
			"FNTA":   []string{"ACVR1"},
			"MAP2K4": []string{"FLNC"},
			"SYNJ1":  []string{"AMPH"},
			"XRN1":   []string{"ALDOA"},
		}
		Expect(Read("test/biogrid.txt", "test/intact.txt", "9606")).To(Equal(expected))
	})
})

var _ = Describe("Parse intact line", func() {
	It("should return source and target when both belong to requested species", func() {
		line := []string{
			"x",
			"x",
			"x",
			"x",
			"psi-mi:pdcd10_human(display_long)|uniprotkb:PDCD10(gene name)|psi-mi:PDCD10(display_short)|uniprotkb:PDCD10(gene name synonym)",
			"psi-mi:stk24_human(display_long)|uniprotkb:STK24(gene name)|psi-mi:STK24(display_short)|uniprotkb:STK24(gene name synonym)",
			"x",
			"x",
			"x",
			"taxid:9606(human)|taxid:9606(Homo sapiens)",
			"taxid:9606(human)|taxid:9606(Homo sapiens)",
		}
		species := "9606"

		actualSource, actualTarget := parseIntactLine(line, species)
		Expect(actualSource).To(Equal("PDCD10"))
		Expect(actualTarget).To(Equal("STK24"))
	})

	It("should return source and target when one belongs to requested species", func() {
		line := []string{
			"x",
			"x",
			"x",
			"x",
			"psi-mi:pdcd10_human(display_long)|uniprotkb:PDCD10(gene name)|psi-mi:PDCD10(display_short)|uniprotkb:PDCD10(gene name synonym)",
			"psi-mi:stk24_human(display_long)|uniprotkb:STK24(gene name)|psi-mi:STK24(display_short)|uniprotkb:STK24(gene name synonym)",
			"x",
			"x",
			"x",
			"taxid:9606(human)|taxid:9606(Homo sapiens)",
			"taxid:10090(human)|taxid:10090(Homo sapiens)",
		}
		species := "9606"

		actualSource, actualTarget := parseIntactLine(line, species)
		Expect(actualSource).To(Equal("PDCD10"))
		Expect(actualTarget).To(Equal("STK24"))
	})

	It("should nil strings when neither belongs to requested species", func() {
		line := []string{
			"x",
			"x",
			"x",
			"x",
			"psi-mi:pdcd10_human(display_long)|uniprotkb:PDCD10(gene name)|psi-mi:PDCD10(display_short)|uniprotkb:PDCD10(gene name synonym)",
			"psi-mi:stk24_human(display_long)|uniprotkb:STK24(gene name)|psi-mi:STK24(display_short)|uniprotkb:STK24(gene name synonym)",
			"x",
			"x",
			"x",
			"taxid:10090(human)|taxid:10090(Homo sapiens)",
			"taxid:10090(human)|taxid:10090(Homo sapiens)",
		}
		species := "9606"

		actualSource, actualTarget := parseIntactLine(line, species)
		Expect(actualSource).To(Equal(""))
		Expect(actualTarget).To(Equal(""))
	})
})
