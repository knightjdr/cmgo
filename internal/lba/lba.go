// Package lba is for localizing prey genes.
package lba

import (
	lbaEnrichment "github.com/knightjdr/cmgo/internal/lba/enrichment"
	lbaLocalize "github.com/knightjdr/cmgo/internal/lba/localize"
)

// Enrichment by GO for LBA.
var Enrichment = lbaEnrichment.Enrichment

// Localize preys from GO enrichment for LBA.
var Localize = lbaLocalize.Localize
