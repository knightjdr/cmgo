// Package countgo count the number of genes with a GO term.
package countgo

import (
	"fmt"
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
)

// Sum counts the number of supplied genes with at least one known GO
// term in the supplied namespace.
func Sum(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	annotations := geneontology.Annotations(options.goAnnotations)
	genes := readGenes(options.genes)

	count, _, _ := countGenesWithAnnotation(genes, (*annotations.Genes)[options.namespace])

	fmt.Println(count, "genes have at least one GO annotation")
}
