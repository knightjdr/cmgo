package uniprot_test

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/uniprot"
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var uniprotText = `
ID   001R_FRG3G              Reviewed;         256 AA.
AC   Q6GZX4;
DT   28-JUN-2011, integrated into UniProtKB/Swiss-Prot.
DT   19-JUL-2004, sequence version 1.
DT   27-SEP-2017, entry version 33.
DE   RecName: Full=Putative transcription factor 001R;
GN   ORFNames=FV3-001R;
OS   Frog virus 3 (isolate Goorha) (FV-3).
OC   Viruses; dsDNA viruses, no RNA stage; Iridoviridae; Alphairidovirinae;
OC   Ranavirus.
OX   NCBI_TaxID=654924;
//
ID   1433B_HUMAN             Reviewed;         246 AA.
AC   P31946; A8K9K2; E1P616;
DT   01-JUL-1993, integrated into UniProtKB/Swiss-Prot.
DT   23-JAN-2007, sequence version 3.
DT   28-MAR-2018, entry version 206.
DE   RecName: Full=14-3-3 protein beta/alpha;
DE   AltName: Full=Protein 1054;
DE   AltName: Full=Protein kinase C inhibitor protein 1;
DE            Short=KCIP-1;
DE   Contains:
DE     RecName: Full=14-3-3 protein beta/alpha, N-terminally processed;
GN   Name=YWHAB;
OS   Homo sapiens (Human).
OC   Eukaryota; Metazoa; Chordata; Craniata; Vertebrata; Euteleostomi;
OC   Mammalia; Eutheria; Euarchontoglires; Primates; Haplorrhini;
OC   Catarrhini; Hominidae; Homo.
OX   NCBI_TaxID=9606;
//
ID   1433E_HUMAN             Reviewed;         255 AA.
AC   P62258; B3KY71; D3DTH5; P29360; P42655; Q4VJB6; Q53XZ5; Q63631;
AC   Q7M4R4;
DT   05-JUL-2004, integrated into UniProtKB/Swiss-Prot.
DT   05-JUL-2004, sequence version 1.
DT   28-MAR-2018, entry version 165.
DE   RecName: Full=14-3-3 protein epsilon;
DE            Short=14-3-3E;
GN   Name=YWHAE;
OS   Homo sapiens (Human).
OC   Eukaryota; Metazoa; Chordata; Craniata; Vertebrata; Euteleostomi;
OC   Mammalia; Eutheria; Euarchontoglires; Primates; Haplorrhini;
OC   Catarrhini; Hominidae; Homo.
OX   NCBI_TaxID=9606;
RN   [1]
//
`

var _ = Describe("Read UniProt data", func() {
	It("should return entries", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/uniprot.txt",
			[]byte(uniprotText),
			0444,
		)

		expected := &uniprot.Entries{
			"P31946": uniprot.Entry{
				Symbol: "YWHAB",
			},
			"P62258": uniprot.Entry{
				Symbol: "YWHAE",
			},
		}
		Expect(uniprot.Read("test/uniprot.txt", 9606)).To(Equal(expected))
	})
})
