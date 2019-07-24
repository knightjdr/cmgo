package matrix

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/mapfunc"
	"github.com/spf13/afero"
)

type cyjs struct {
	Elements elements `json:"elements"`
}

type edge struct {
	Data edgeData `json:"data"`
}

type edgeData struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Source string  `json:"source"`
	Target string  `json:"target"`
	Weight float64 `json:"weight"`
}

type elements struct {
	Nodes []node `json:"nodes"`
	Edges []edge `json:"edges"`
}

type node struct {
	Data     nodeData `json:"data"`
	Position position `json:"position"`
}

type nodeData struct {
	Color       string `json:"color"`
	Compartment string `json:"compartment"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

type position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func compartmentOrder(compartments []string) map[string]int {
	order := make(map[string]int, len(compartments))
	for i, compartment := range compartments {
		order[compartment] = i
	}
	return order
}

func fillColor(index int, name string, colors []string) string {
	if name != "" {
		return colors[index+1]
	}
	return colors[0]
}

func firstLocalization(localizations map[string]string) (string, string) {
	ids := mapfunc.KeysStringString((localizations))
	sort.Strings(ids)
	if len(ids) > 0 {
		return ids[0], localizations[ids[0]]
	}
	return "", ""
}

func writeJSON(
	matrix [][]float64,
	genes []string,
	cutoff float64,
	nodeLocalizations map[string]map[string]string,
	possibleLocalizations []string,
	colors []string,
	outfile string,
) {
	localizationOrder := compartmentOrder(possibleLocalizations)
	fileData := &cyjs{}
	for i, gene := range genes {
		compartmentID, compartmentName := firstLocalization(nodeLocalizations[gene])
		geneNodeData := nodeData{
			Color:       fillColor(localizationOrder[compartmentID], compartmentName, colors),
			Compartment: compartmentName,
			ID:          gene,
			Name:        gene,
		}
		genePosition := position{
			X: 1,
			Y: 1,
		}
		fileData.Elements.Nodes = append(fileData.Elements.Nodes, node{Data: geneNodeData, Position: genePosition})
		for j := i + 1; j < len(matrix); j++ {
			weight := float64(0)
			for k := range matrix[j] {
				if matrix[i][k] > weight && matrix[i][k] >= cutoff && matrix[j][k] >= cutoff {
					weight = matrix[i][k]
				}
			}

			if weight != 0 {
				edgeNodeData := edgeData{
					ID:     fmt.Sprintf("%s-%s", genes[i], genes[j]),
					Name:   fmt.Sprintf("%s (interacts with) %s", genes[i], genes[j]),
					Source: genes[i],
					Target: genes[j],
					Weight: weight,
				}
				fileData.Elements.Edges = append(fileData.Elements.Edges, edge{Data: edgeNodeData})
			}
		}
	}

	bytes, err := json.MarshalIndent(fileData, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	afero.WriteFile(fs.Instance, outfile, bytes, 0644)
}
