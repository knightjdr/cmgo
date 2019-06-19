// Package localize localizes preys using LBA
package localize

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/database"
	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	"github.com/knightjdr/cmgo/pkg/gene"
)

// Localize localizes preys from a SAINT file using LBA
func Localize(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	saintData := saint.Read(options.saintFile, options.fdr, options.minBaits)
	preys := uniquePreys(saintData)

	// Read database and create a mapping from Refseq to gene name,
	// gene ID and UniProt ID for significant preys
	databaseData := database.Read(options.database, false)
	refseqMapping, geneIDs := mapRefseq(databaseData, preys)
	entrezToUniprotMap := gene.MapIDs(geneIDs, "Entrez", "UniProt", "")
	addUniprotIDs(&refseqMapping, entrezToUniprotMap)

	// Read GO terms.
	/* goAnnotations := */
	geneontology.Annotations(options.goAnnotations)
	goHierarchy := geneontology.OBO(options.goHierarchy)
	goHierarchy.GetChildren(options.namespace)
	goHierarchy.GetParents(options.namespace)

	/* preyAssociations := countAssociations(saintData) */
}
