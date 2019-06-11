package geneontology

import (
	"testing"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var goOboText = `format-version: 1.2
data-version: releases/2019-02-13
subsetdef: gocheck_do_not_annotate "Term not to be used for direct annotation"
subsetdef: gocheck_do_not_manually_annotate "Term not to be used for direct manual annotation"
ontology: go

[Term]
id: GO:0000001
name: mitochondrion inheritance
namespace: biological_process
def: "" [GOC:mcc, PMID:10873824, PMID:11389764]
synonym: "mitochondrial inheritance" EXACT []
is_a: GO:0048308 ! organelle inheritance
is_a: GO:0048311 ! mitochondrion distribution

[Term]
id: GO:0000002
name: mitochondrial genome maintenance
namespace: biological_process
def: "" [GOC:ai, GOC:vw]
is_a: GO:0007005 ! mitochondrion organization

[Term]
id: GO:0000005
name: obsolete ribosomal chaperone activity
namespace: molecular_function
is_obsolete: true
consider: GO:0042254
consider: GO:0044183
consider: GO:0051082

id: GO:0005652
name: nuclear lamina
namespace: cellular_component
def: "" [ISBN:0198506732, ISBN:0716731363]
xref: NIF_Subcellular:sao1455996588
xref: Wikipedia:Nuclear_lamina
is_a: GO:0044428 ! nuclear part
relationship: part_of GO:0034399 ! nuclear periphery

[Term]
id: GO:0005654
name: nucleoplasm
namespace: cellular_component
def: "" [GOC:ma, ISBN:0124325653]
subset: goslim_chembl
subset: goslim_generic
subset: goslim_plant
xref: NIF_Subcellular:sao661522542
xref: Wikipedia:Nucleoplasm
is_a: GO:0044428 ! nuclear part
relationship: part_of GO:0031981 ! nuclear lumen

`

func TestOBO(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/go.obo",
		[]byte(goOboText),
		0444,
	)

	// TEST
	wanted := &GOhierarchy{
		"BP": map[string]*GOterm{
			"GO:0000001": &GOterm{
				DirectParents: []string{"GO:0048308", "GO:0048311"},
				Name:          "mitochondrion inheritance",
				Synonyms: GOsynonym{
					"Broad":   []string{},
					"Exact":   []string{"mitochondrial inheritance"},
					"Narrow":  []string{},
					"Related": []string{},
				},
			},
			"GO:0000002": &GOterm{
				DirectParents: []string{"GO:0007005"},
				Name:          "mitochondrial genome maintenance",
				Synonyms: GOsynonym{
					"Broad":   []string{},
					"Exact":   []string{},
					"Narrow":  []string{},
					"Related": []string{},
				},
			},
		},
		"CC": map[string]*GOterm{
			"GO:0005652": &GOterm{
				DirectParents: []string{"GO:0044428", "GO:0034399"},
				Name:          "nuclear lamina",
				Synonyms: GOsynonym{
					"Broad":   []string{},
					"Exact":   []string{},
					"Narrow":  []string{},
					"Related": []string{},
				},
			},
			"GO:0005654": &GOterm{
				DirectParents: []string{"GO:0044428", "GO:0031981"},
				Name:          "nucleoplasm",
				Synonyms: GOsynonym{
					"Broad":   []string{},
					"Exact":   []string{},
					"Narrow":  []string{},
					"Related": []string{},
				},
			},
		},
		"MF": map[string]*GOterm{},
	}
	assert.Equal(t, wanted, OBO("test/go.obo"), "Should read go.obo hierarchy")
}
