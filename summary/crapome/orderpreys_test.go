package crapome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderPreys(t *testing.T) {
	data := map[string]map[string]int{
		"B": map[string]int{
			"128_468": 4,
			"128_737": 8,
		},
		"A": map[string]int{
			"128_468": 5,
			"128_737": 10,
		},
		"C": map[string]int{
			"128_590": 2,
		},
		"D": map[string]int{
			"128_590": 3,
		},
	}
	preyMap := map[string]string{
		"A": "preyA",
		"B": "preyB",
		"C": "preyC",
		"D": "preyD",
		"E": "preyE",
		"F": "preyF",
	}
	wanted := []string{"A", "B", "C", "D"}
	assert.Equal(t, wanted, orderPreys(data, preyMap), "Should order preys to output alphabetically")
}
