package organelle

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
)

// Compartments contains the name and list of proteins belonging to each compartment
// to compare.
type Compartments []struct {
	Name     string   `json:"name"`
	Proteins []string `json:"proteins"`
}

// ReadCompartments reads JSON file with compartments and their proteins
func ReadCompartments(filename string) Compartments {
	byteValue, err := afero.ReadFile(fs.Instance, filename)
	if err != nil {
		fmt.Println(filename)
		log.Fatalln(err)
	}
	var compartments Compartments
	json.Unmarshal(byteValue, &compartments)

	return compartments
}
