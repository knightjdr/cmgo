package prediction

import (
	customMath "github.com/knightjdr/cmgo/pkg/math"
	"github.com/knightjdr/cmgo/pkg/slice"
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
		predictedCompartment := predictions[gene]
		for _, domain := range domains {
			if isDomainSupportive(domain, compartmentDomains, predictedCompartment) {
				(*scores)[gene].supportingDomains = append((*scores)[gene].supportingDomains, domain)
			}
			if isDomainConflicting(domain, compartmentDomains, predictedCompartment) {
				(*scores)[gene].conflictingDomains = append((*scores)[gene].conflictingDomains, domain)
			}
		}
		(*scores)[gene].score = calculateDomainScore((*scores)[gene])
	}
	return scores
}

func isDomainSupportive(domain string, compartmentDomains map[int][]string, predictedPreyCompartment int) bool {
	if slice.ContainsString(domain, compartmentDomains[predictedPreyCompartment]) {
		return true
	}
	return false
}

func isDomainConflicting(domain string, compartmentDomains map[int][]string, predictedPreyCompartment int) bool {
	for compartment := range compartmentDomains {
		if compartment != predictedPreyCompartment && isDomainSupportive(domain, compartmentDomains, compartment) {
			return true
		}
	}
	return false
}

func calculateDomainScore(scoreComponents *domainScoreComponents) float64 {
	score := float64(0)
	if scoreComponents.totalDomains > 0 {
		supporting := float64(len(scoreComponents.supportingDomains))
		total := float64(scoreComponents.totalDomains)
		score = customMath.Round(supporting/total, 0.00001)
	}
	return score
}
