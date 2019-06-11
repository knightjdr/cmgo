package main

import (
	"errors"
	"log"

	"github.com/knightjdr/cmgo/internal/analysis/dbgenes"
	"github.com/knightjdr/cmgo/internal/assessment/bait/gradient"
	"github.com/knightjdr/cmgo/internal/assessment/localization/nmfsafe"
	"github.com/knightjdr/cmgo/internal/enrichment/heatmap"
	"github.com/knightjdr/cmgo/internal/network/svg"
	"github.com/knightjdr/cmgo/internal/nmf/subset"
	"github.com/knightjdr/cmgo/internal/organelle/overlap"
	"github.com/knightjdr/cmgo/internal/organelle/shared"
	"github.com/knightjdr/cmgo/internal/summary/crapome"
	"github.com/knightjdr/cmgo/internal/summary/notsignificant"
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
