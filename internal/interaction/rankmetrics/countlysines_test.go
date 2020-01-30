package rankmetrics

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var fastaText = `>NP_000000001|gn|GENEa:00001| gene A [Homo sapiens]
MATARTFGPEREAEPAKEARVVGSELVDTYTVYIIQVTDGSHEWTVKHRYSDFHDLHEKLVAERKIDKNL
LPPKKIIGKNSRSLVEKREKDLEVYLQKLLAAFPGVTPRVLAHFLHFHFYEINGITAALAEELFEKGEQL
>NP_000000002|gn|GENEb:00002| gene b [Homo sapiens]
LGAGEVFAIGPLQLYAVTEQLQQGKPTCASGDAKTDLGHILDFTCRLKYLKVSGTEGPFGTSNIQEQLLP
FDLSIFKSLHQVEISHCDAKHIRGLVASKPTLATLSVRFSATSMKEVLVPEASEFDEWEPEGTTLEGPVT
>NP_000000003|gn|GENEc:00003| gene c [Homo sapiens]
AVIPTWQALTTLDLSHNSISEIDESVKLIPKIEFLDLSHNGLLVVDNLQHLYNLVHLDLSYNKLSSLEGL
HTKLGNIKTLNLAGNLLESLSGLHKLYSLVNLDLRDNRIEQMEEVRSIGSLPCLEHVSLLNNPLSIIPDY
`

var _ = Describe("Count lysines", func() {
	It("should count the number of lysines per gene", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/database.fasta",
			[]byte(fastaText),
			0444,
		)

		expected := map[string]int{
			"GENEa": 12,
			"GENEb": 8,
			"GENEc": 6,
		}
		Expect(countLysines("test/database.fasta")).To(Equal(expected))
	})
})
