package pfam_test

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/pfam"
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var domainText = `#<seq id> <alignment start> <alignment end> <envelope start> <envelope end> <hmm acc> <hmm name> <type> <hmm start> <hmm end> <hmm length> <bit score> <E-value> <clan>
A0A024QZ18	69	147	66	147	PF00595	PDZ	PfamLive::Result::SequenceOntology=HASH(0x8b8beb0)	4	82	82	51.30	2.8e-10	CL0466
A0A024QZ33	5	123	4	123	PF09745	DUF2040	PfamLive::Result::SequenceOntology=HASH(0xfaeac30)	2	121	121	124.70	5.1e-33	No_clan
O95996	1621	1642	1621	1642	PF05924	SAMP	PfamLive::Result::SequenceOntology=HASH(0xee9f888)	1	22	22	26.00	0.0099	No_clan
O95996	1786	2116	1786	2118	PF05956	APC_basic	PfamLive::Result::SequenceOntology=HASH(0xeeb8100)	1	339	342	369.20	6.1e-107	No_clan
`

var _ = Describe("Read domains", func() {
	It("should return domains from a file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/domains.txt",
			[]byte(domainText),
			0444,
		)

		expected := pfam.Domains{
			"A0A024QZ18": []pfam.Domain{
				pfam.Domain{End: 147, Name: "PDZ", Start: 69},
			},
			"A0A024QZ33": []pfam.Domain{
				pfam.Domain{End: 123, Name: "DUF2040", Start: 5},
			},
			"O95996": []pfam.Domain{
				pfam.Domain{End: 1642, Name: "SAMP", Start: 1621},
				pfam.Domain{End: 2116, Name: "APC_basic", Start: 1786},
			},
		}
		Expect(pfam.ReadDomains("test/domains.txt")).To(Equal(expected))
	})
})
