package prediction

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate domain score component", func() {
	It("should return score and domain lists for each prey", func() {
		compartmentDomains := map[int][]string{
			1: []string{"domain1", "domain2"},
			2: []string{"domain2", "domain3", "domain4"},
		}
		geneDomains := map[string][]string{
			"prey1": []string{"domain1", "domain1", "domain3", "domain5"},
			"prey2": []string{"domain2", "domain2", "domain4", "domain5"},
			"prey3": []string{"domain2", "domain3", "domain4", "domain5"},
			"prey4": []string{"domain1"},
		}
		predictions := map[string]int{
			"prey1": 1,
			"prey2": 1,
			"prey3": 2,
			"prey4": 1,
		}

		expected := preyDomainScore{
			"prey1": &domainScoreComponents{
				conflictingDomains: []string{"domain3"},
				score:              0.5,
				supportingDomains:  []string{"domain1", "domain1"},
				totalDomains:       4,
			},
			"prey2": &domainScoreComponents{
				conflictingDomains: []string{"domain2", "domain2", "domain4"},
				score:              0.5,
				supportingDomains:  []string{"domain2", "domain2"},
				totalDomains:       4,
			},
			"prey3": &domainScoreComponents{
				conflictingDomains: []string{"domain2"},
				score:              0.75,
				supportingDomains:  []string{"domain2", "domain3", "domain4"},
				totalDomains:       4,
			},
			"prey4": &domainScoreComponents{
				conflictingDomains: []string{},
				score:              1,
				supportingDomains:  []string{"domain1"},
				totalDomains:       1,
			},
		}
		actual := calculateDomainComponentScore(geneDomains, compartmentDomains, predictions)
		for prey, scoreComponent := range *actual {
			Expect(scoreComponent.conflictingDomains).To(Equal(expected[prey].conflictingDomains), "should have conflicting domains")
			Expect(scoreComponent.score).To(BeNumerically("~", expected[prey].score, 0.00001))
			Expect(scoreComponent.supportingDomains).To(Equal(expected[prey].supportingDomains), "should have supporting domains")
			Expect(scoreComponent.totalDomains).To(Equal(expected[prey].totalDomains), "should have total domain number")
		}
	})
})
