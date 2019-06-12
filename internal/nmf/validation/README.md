# Module name: `nmf-validation`

> Takes assigned localization(s) for each NMF rank and tests...

| parameters | definition | default |
|------------|------------|---------|
| basisMatrix | | |
| maxGenesPerRank | maximum number of genes to use per rank for enrichments | 100 |
| minRankValue | a prey must have an NMF value at or above this value to be used for enrichment | 0.25 |
| outFile | output file name | basis-subset.svg |
| withinRankMax | if a prey has an NMF value within this % of max in its non-primary rank, it can be used for enrichment | 0.75 |

## Example file formats

### basisMatrix

### Output
* 