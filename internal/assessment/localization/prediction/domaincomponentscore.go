package prediction

import (
	"fmt"
	"github.com/knightjdr/cmgo/pkg/math"
	"github.com/knightjdr/cmgo/pkg/slice"
	"strconv"
)

type preyDomainScore map[string]*domainScoreComponents

type domainScoreComponents struct {
	conflictingDomains []string
	score              float64
	supportingDomains  []string
	totalDomains       int
}

func calculateDomainComponentScore(geneDomains map[string][]string, compartmentDomains map[int][]string, predictions map[string]int) *preyDomainScore {
	scores := &preyDomainScore{}

	for gene, domains := range geneDomains {
		(*scores)[gene] = &domainScoreComponents{
			conflictingDomains: make([]string, 0),
			supportingDomains:  make([]string, 0),
			totalDomains:       len(domains),
		}
		compartment := predictions[gene]
		for _, domain := range domains {
			if slice.ContainsString(domain, compartmentDomains[compartment]) {
				(*scores)[gene].supportingDomains = append((*scores)[gene].supportingDomains, domain)
			}
		}
		(*scores)[gene].score = calculateDomainScore((*scores)[gene].supportingDomains, (*scores)[gene].totalDomains)
	}
	return scores
}

func calculateDomainScore(supporting []string, total int) float64 {
	score := float64(0)
	if total > 0 {
		score = math.Round(float64(len(supporting))/float64(total), 0.00001)
	}
	scoreString := fmt.Sprintf("%0.5f", score)
	scoreFloat, _ := strconv.ParseFloat(scoreString, 64)
	return scoreFloat
}
