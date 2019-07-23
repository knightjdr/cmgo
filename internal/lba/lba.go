// Package lba is for localizing prey genes.
package lba

import (
	lbaCorrelation "github.com/knightjdr/cmgo/internal/lba/correlation"
	lbaEnrichment "github.com/knightjdr/cmgo/internal/lba/enrichment"
	lbaLocalize "github.com/knightjdr/cmgo/internal/lba/localize"
)

// Correlation creates a network using LBA profile correlation.
var Correlation = lbaCorrelation.Network

// Enrichment by GO for LBA.
var Enrichment = lbaEnrichment.Enrichment

// Localize preys from GO enrichment for LBA.
var Localize = lbaLocalize.Localize
