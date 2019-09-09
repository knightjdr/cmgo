// Package dbgenes generates a list of genes in sequence database
package dbgenes

import (
	"log"
)

// List reads FASTA database and outputs gene names.
func List(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	genes := readList(options.ncbigene)
	writeGenes(genes, options.outFile)
}
