// Package interactions reads interactions from BioGRID and/or Intact
package interactions

import (
	"regexp"
)

var reGene = regexp.MustCompile(`uniprotkb:([^\(]+)\(gene name\)`)
var reTaxon = regexp.MustCompile(`^taxid:([-\d]+)\(`)

// Read known interactions from BioGRID and IntAct
func Read(biogridFilename, intactFilename, species string) map[string][]string {
	biogridInteractors := readFile(biogridFilename, "", parseBiogridLine)
	intactInteractors := readFile(intactFilename, species, parseIntactLine)

	return mergeInteractors(biogridInteractors, intactInteractors)
}

func parseBiogridLine(line []string, species string) (string, string) {
	return line[7], line[8]
}

func parseIntactLine(line []string, species string) (string, string) {
	sourceSpecies := parseValueWithRegex(line[9], reTaxon)
	targetSpecies := parseValueWithRegex(line[10], reTaxon)
	if isCorrectSpecies(species, sourceSpecies, targetSpecies) {
		sourceGene := parseValueWithRegex(line[4], reGene)
		targetGene := parseValueWithRegex(line[5], reGene)
		return sourceGene, targetGene
	}
	return "", ""
}

func parseValueWithRegex(str string, re *regexp.Regexp) (parsedValue string) {
	matches := re.FindStringSubmatch(str)
	if len(matches) > 0 {
		parsedValue = matches[1]
	}
	return
}

func isCorrectSpecies(species, sourceSpecies, targetSpecies string) bool {
	if sourceSpecies == species && targetSpecies != "" {
		return true
	}
	if targetSpecies == species && sourceSpecies != "" {
		return true
	}
	return false
}
