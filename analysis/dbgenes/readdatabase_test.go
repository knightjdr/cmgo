package dbgenes

import (
	"sort"
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var databaseText = `>NP_001263222.1|gn|NISCH:11188| nischarin isoform 2 [Homo sapiens]
MATARTFGPEREAEPAKEARVVGSELVDTYTVYIIQVTDGSHEWTVKHRYSDFHDLHEKLVAERKIDKNL
>NP_001263218.1|gn|PRKAR1A:5573| cAMP-dependent protein kinase type I-alpha regulatory subunit isoform a [Homo sapiens]
MESGSTAASEEARSLRECELYVQKHNIQALLKDSIVQLCTARPERPMAFLREYFERLEKEEAKQIQNLQK
AGTRTDSREDEISPPPPNPVVKGRRRRGAISAEVYTEEDAASYVRKVIPKDYKTMAALAKAIEKNVLFSH
>NP_001263215.1|gn|BBX:56987| HMG box transcription factor BBX isoform 3 [Homo sapiens]
MKGSNRNKDHSAEGEGVGKRPKRKCLQWHPLLAKKLLDFSEEEEEEDEEEDIDKVQLLGADGLEQDVGET
EDDESPEQRARRPMNAFLLFCKRHRSLVRQEHPRLDNRGATKILADWWAVLDPKEKQKYTDMAKEYKDAF
>NP_001263215.1|gn|BBX:56987| HMG box transcription factor BBX isoform 3 [Homo sapiens]
MKGSNRNKDHSAEGEGVGKRPKRKCLQWHPLLAKKLLDFSEEEEEEDEEEDIDKVQLLGADGLEQDVGET
EDDESPEQRARRPMNAFLLFCKRHRSLVRQEHPRLDNRGATKILADWWAVLDPKEKQKYTDMAKEYKDAF
`

func TestReadDatabase(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/db.fasta",
		[]byte(databaseText),
		0444,
	)

	// TEST1: only filter by FDR
	wanted := []string{"BBX", "NISCH", "PRKAR1A"}
	results := readDatabase("test/db.fasta")
	sort.Strings(results)
	assert.Equal(t, wanted, results, "Should return unique gene names from database")
}
