package rankmetrics

import (
	"encoding/json"
	"log"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

type gixData []gixEntry

type gixEntry struct {
	Expression gixExpression `json:"protein-expression"`
	Uniprot    []string
}

type gixExpression struct {
	Cells gixHek293 `json:"cells"`
}

type gixHek293 struct {
	Hek293 gixIntensity `json:"HEK-293"`
}

type gixIntensity struct {
	Intensity float64
}

func readExpression(data analysis) map[string]float64 {
	byteValue, err := afero.ReadFile(fs.Instance, data.parameters.gixDB)
	if err != nil {
		log.Fatalln(err)
	}

	var expressionData gixData
	json.Unmarshal(byteValue, &expressionData)

	expression := make(map[string]float64)
	for _, entry := range expressionData {
		for _, uniprot := range entry.Uniprot {
			if _, ok := data.uniprotMapping[uniprot]; ok {
				gene := data.uniprotMapping[uniprot]
				expression[gene] = entry.Expression.Cells.Hek293.Intensity
			}
		}
	}

	return expression
}
