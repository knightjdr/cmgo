package tsnecytoscape

import (
	"encoding/json"
	"log"
	"sort"

	"github.com/knightjdr/cmgo/internal/pkg/read/tsne"
	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/mapfunc"
	"github.com/spf13/afero"
)

type cyjs struct {
	Elements elements `json:"elements"`
}

type data struct {
	Color       string `json:"color"`
	Compartment string `json:"compartment"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

type edge struct{}

type elements struct {
	Nodes []node `json:"nodes"`
	Edges []edge `json:"edges"`
}

type node struct {
	Data     data     `json:"data"`
	Position position `json:"position"`
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
	nodeCoordinates map[string]tsne.Coordinate,
	nodeLocalizations map[string]map[string]string,
	possibleLocalizations []string,
	colors []string,
	transformation map[string]float64,
	outfile string,
) {
	geneOrder := make([]string, len(nodeCoordinates))
	i := 0
	for gene := range nodeCoordinates {
		geneOrder[i] = gene
		i++
	}
	sort.Strings(geneOrder)

	localizationOrder := compartmentOrder(possibleLocalizations)
	fileData := &cyjs{}
	for _, gene := range geneOrder {
		coordinates := nodeCoordinates[gene]
		compartmentID, compartmentName := firstLocalization(nodeLocalizations[gene])
		geneData := data{
			Color:       fillColor(localizationOrder[compartmentID], compartmentName, colors),
			Compartment: compartmentName,
			ID:          gene,
			Name:        gene,
		}
		genePosition := position{
			X: transformation["scale"] * (coordinates.X + transformation["translateX"]),
			Y: transformation["scale"] * (coordinates.Y + transformation["translateY"]),
		}
		fileData.Elements.Nodes = append(fileData.Elements.Nodes, node{Data: geneData, Position: genePosition})
	}
	fileData.Elements.Edges = []edge{}

	bytes, err := json.MarshalIndent(fileData, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	afero.WriteFile(fs.Instance, outfile, bytes, 0644)
}
