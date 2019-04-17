package main

import (
	"errors"
	"log"

	"github.com/knightjdr/cmgo/organelle/overlap"
	"github.com/knightjdr/cmgo/organelle/shared"
)

func main() {
	options, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	if options["module"] == "organelle-overlap" {
		overlap.Metrics(options)
	} else if options["module"] == "organelle-sharedregion" {
		shared.Region(options)
	} else {
		log.Fatalln(errors.New("Unknown analysis module"))
	}
}
