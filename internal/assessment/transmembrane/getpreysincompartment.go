package transmembrane

import (
	"strconv"
)

func getPreysInCompartment(preysPerRank map[int][]string, ranks []string) []string {
	preys := make([]string, 0)
	for _, rank := range ranks {
		integerRank, _ := strconv.Atoi(rank)
		preys = append(preys, preysPerRank[integerRank]...)
	}

	return preys
}
