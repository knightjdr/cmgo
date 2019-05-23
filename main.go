package main

import (
	"errors"
	"log"

	"github.com/knightjdr/cmgo/analysis/dbgenes"
	"github.com/knightjdr/cmgo/assessment/bait/gradient"
	"github.com/knightjdr/cmgo/assessment/localization/nmfsafe"
	"github.com/knightjdr/cmgo/enrichment/heatmap"
	"github.com/knightjdr/cmgo/network/svg"
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

	switch module := options["module"]; module {
	case "analysis-dbgenes":
		dbgenes.List(options)
	case "bait-gradient":
		gradient.Draw(options)
	case "enrichment-heatmap":
		heatmap.Region(options)
	case "network-svg":
		svg.Draw(options)
	case "nmf-subset":
		subset.NMF(options)
	case "nmf-v-safe":
		nmfsafe.Concordance(options)
	case "organelle-overlap":
		overlap.Metrics(options)
	case "organelle-sharedregion":
		shared.Region(options)
	case "summary-crapome":
		crapome.Matrix(options)
	case "summary-notsignificant":
		notsignificant.List(options)
	default:
		log.Fatalln(errors.New("Unknown analysis module"))
	}
}
