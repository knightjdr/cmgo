package main

import (
	"log"

	"github.com/knightjdr/cmgo/organelle/overlap"
)

func main() {
	options, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	if options["module"] == "organelle-overlap" {
		overlap.Metrics(options)
	}
}
