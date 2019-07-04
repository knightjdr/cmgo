// Package localize localizes preys using LBA
package localize

import (
	"fmt"
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/database"
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

	baitsPerPrey, preysPerBait := associations(saintData)
	_, topPreysPerPrey := topPreyPartners(baitsPerPrey, preysPerBait, options.preyLimit, options.minFC)
	enrichment := profile(topPreysPerPrey, preys)
	for gene, terms := range enrichment {
		for _, term := range terms {
			fmt.Println(gene, term.Recall, term.Precision, term.Name, term.Pvalue)
		}
	}
}
