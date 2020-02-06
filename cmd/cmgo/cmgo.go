package main

import (
	"errors"
	"log"

	"github.com/knightjdr/cmgo/internal/analysis/dbgenes"
	"github.com/knightjdr/cmgo/internal/assessment/bait/gradient"
	"github.com/knightjdr/cmgo/internal/assessment/controls/preys"
	"github.com/knightjdr/cmgo/internal/assessment/countgo"
	"github.com/knightjdr/cmgo/internal/assessment/hydropathy"
	assessLocalization "github.com/knightjdr/cmgo/internal/assessment/localization/evaluate"
	"github.com/knightjdr/cmgo/internal/assessment/localization/nmfsafe"
	"github.com/knightjdr/cmgo/internal/assessment/localization/prediction"
	"github.com/knightjdr/cmgo/internal/assessment/localization/recovered"
	"github.com/knightjdr/cmgo/internal/assessment/transmembrane"
	"github.com/knightjdr/cmgo/internal/enrichment/genes"
	"github.com/knightjdr/cmgo/internal/enrichment/heatmap"
	"github.com/knightjdr/cmgo/internal/interaction/knownbyrank"
	"github.com/knightjdr/cmgo/internal/interaction/rankaverage"
	"github.com/knightjdr/cmgo/internal/interaction/rankmetrics"
	"github.com/knightjdr/cmgo/internal/lba"
	"github.com/knightjdr/cmgo/internal/network/correlation"
	"github.com/knightjdr/cmgo/internal/network/matrix"
	"github.com/knightjdr/cmgo/internal/network/svg"
	"github.com/knightjdr/cmgo/internal/network/tsnecytoscape"
	"github.com/knightjdr/cmgo/internal/nmf/moonlighting"
	"github.com/knightjdr/cmgo/internal/nmf/robustness"
	"github.com/knightjdr/cmgo/internal/nmf/subset"
	"github.com/knightjdr/cmgo/internal/nmf/uv"
	"github.com/knightjdr/cmgo/internal/organelle/isolation"
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
	case "assessment-compartment-recovered":
		recovered.AssessCompartment(options)
	case "assessment-countgo":
		countgo.Sum(options)
	case "assessment-hydropathy":
		hydropathy.Assess(options)
	case "assessment-localization":
		assessLocalization.Evaluate(options)
	case "assessment-prediction":
		prediction.Score(options)
	case "assessment-transmembrane":
		transmembrane.Orientation(options)
	case "bait-gradient":
		gradient.Draw(options)
	case "control-preys":
		preys.Assess(options)
	case "enrichment-genes":
		genes.Enrich(options)
	case "enrichment-heatmap":
		heatmap.Region(options)
	case "interaction-knownbyrank":
		knownbyrank.CalculateKnownByRank(options)
	case "interaction-rankaverage":
		rankaverage.CalculateRankAverages(options)
	case "interaction-rankmetrics":
		rankmetrics.Calculate(options)
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
	case "nmf-moonlighting":
		moonlighting.Calculate(options)
	case "nmf-robustness":
		robustness.Evaluate(options)
	case "nmf-subset":
		subset.NMF(options)
	case "nmf-uv":
		uv.Assess(options)
	case "nmf-v-safe":
		nmfsafe.Concordance(options)
	case "organelle-overlap":
		overlap.Metrics(options)
	case "organelle-isolation":
		isolation.Calculate(options)
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
