package main

import (
	"errors"
	"log"

	"github.com/knightjdr/cmgo/enrichment/heatmap"
	"github.com/knightjdr/cmgo/nmf/subset"
	"github.com/knightjdr/cmgo/organelle/overlap"
	"github.com/knightjdr/cmgo/organelle/shared"
	"github.com/knightjdr/cmgo/summary/crapome"
	"github.com/knightjdr/cmgo/summary/notsignificant"
)

func main() {
	options, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	if options["module"] == "enrichment-heatmap" {
		heatmap.Region(options)
	} else if options["module"] == "nmf-subset" {
		subset.NMF(options)
	} else if options["module"] == "organelle-overlap" {
		overlap.Metrics(options)
	} else if options["module"] == "organelle-sharedregion" {
		shared.Region(options)
	} else if options["module"] == "summary-crapome" {
		crapome.Matrix(options)
	} else if options["module"] == "summary-notsignificant" {
		notsignificant.List(options)
	} else {
		log.Fatalln(errors.New("Unknown analysis module"))
	}
}
