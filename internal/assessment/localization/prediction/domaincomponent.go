package prediction

import (
	"sort"
	"strconv"

	"github.com/knightjdr/cmgo/internal/pkg/read/list"
	"github.com/knightjdr/cmgo/internal/pkg/read/pfam"
)

func calculateDomainComponent(options parameters, inputFiles fileContent) *preyDomainScore {
	compartmentDomains := readCompartmentDomains(options.domainsPerCompartment)
	geneDomains := readGeneDomains(options.domainsPerGene, inputFiles)
	return calculateDomainComponentScore(geneDomains, compartmentDomains, inputFiles.predictions)
}

func readCompartmentDomains(filename string) map[int][]string {
	fileContent := list.CSV(filename, '\t')

	domains := make(map[int][]string, 0)
	for _, entry := range fileContent {
		compartment, _ := strconv.Atoi(entry["rank"])
		if _, ok := domains[compartment]; !ok {
			domains[compartment] = make([]string, 0)
		}
		domains[compartment] = append(domains[compartment], entry["term"])
	}

	return domains
}

func readGeneDomains(domainFile string, inputFiles fileContent) map[string][]string {
	pfamDomains := pfam.ReadDomains(domainFile)

	domains := make(map[string][]string, 0)
	for gene := range inputFiles.predictions {
		domains[gene] = make([]string, 0)
		if _, ok := inputFiles.geneToUniProt[gene]; ok {
			id := inputFiles.geneToUniProt[gene]
			for _, domain := range pfamDomains[id] {
				domains[gene] = append(domains[gene], domain.Name)
			}
		}
		sort.Strings(domains[gene])
	}

	return domains
}
