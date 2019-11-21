package main

import (
	"errors"
	"log"

	"github.com/knightjdr/cmgo/internal/analysis/dbgenes"
	"github.com/knightjdr/cmgo/internal/assessment/bait/gradient"
	"github.com/knightjdr/cmgo/internal/assessment/hydropathy"
	assessLocalization "github.com/knightjdr/cmgo/internal/assessment/localization/evaluate"
	"github.com/knightjdr/cmgo/internal/assessment/localization/nmfsafe"
	"github.com/knightjdr/cmgo/internal/enrichment/genes"
	"github.com/knightjdr/cmgo/internal/enrichment/heatmap"
	"github.com/knightjdr/cmgo/internal/interaction/knownbyrank"
	"github.com/knightjdr/cmgo/internal/lba"
	"github.com/knightjdr/cmgo/internal/network/correlation"
	"github.com/knightjdr/cmgo/internal/network/matrix"
	"github.com/knightjdr/cmgo/internal/network/svg"
	"github.com/knightjdr/cmgo/internal/network/tsnecytoscape"
	"github.com/knightjdr/cmgo/internal/nmf/robustness"
	"github.com/knightjdr/cmgo/internal/nmf/subset"
	"github.com/knightjdr/cmgo/internal/nmf/uv"
	"github.com/knightjdr/cmgo/internal/organelle/overlap"
	"github.com/knightjdr/cmgo/internal/organelle/shared"
	preypreySubset "github.com/knightjdr/cmgo/internal/preyprey/subset"
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
	case "assessment-hydropathy":
		hydropathy.Assess(options)
	case "assessment-localization":
		assessLocalization.Evaluate(options)
	case "bait-gradient":
		gradient.Draw(options)
	case "enrichment-genes":
		genes.Enrich(options)
	case "enrichment-heatmap":
		heatmap.Region(options)
	case "interaction-knownbyrank":
		knownbyrank.CalculateKnownByRank(options)
	case "lba-enrichment":
		lba.Enrichment(options)
	case "lba-localize":
		lba.Localize(options)
	case "network-correlation":
		correlation.Network(options)
	case "network-matrix":
		matrix.Network(options)
	case "network-svg":
		svg.Draw(options)
	case "network-tsnecytoscape":
		tsnecytoscape.Create(options)
	case "nmf-robustness":
		robustness.Evaluate(options)
	case "nmf-uv":
		uv.Assess(options)
	case "nmf-subset":
		subset.NMF(options)
	case "nmf-v-safe":
		nmfsafe.Concordance(options)
	case "organelle-overlap":
		overlap.Metrics(options)
	case "organelle-sharedregion":
		shared.Region(options)
	case "preyprey-subset":
		preypreySubset.Heatmap(options)
	case "summary-crapome":
		crapome.Matrix(options)
	case "summary-notsignificant":
		notsignificant.List(options)
	default:
		log.Fatalln(errors.New("Unknown analysis module"))
	}
}
