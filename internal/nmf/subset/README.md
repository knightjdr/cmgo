# Module name: `nmf-subset`

> Subset an NMF basis (prey) matrix to only include preys enriched in specified compartments.

| parameters | definition | default |
|------------|------------|---------|
| abundanceCap | the value to cap cell color on the heat map | 1 |
| basisMatrix | NMF basis matrix | |
| clusteringMethod | the hierarchical clustering method | complete |
| distanceMetric | the distance metric | euclidean |
| minAbundance | the minimum abundance for limiting the color on the heat map | 0 |
| minNMFScore | the minimum score required in any of the specified ranks to include a row | 0 |
| outFile | output file name | basis-subset.svg |
| ranks1 | ranks defining first NMF compartment | |
| ranks2 | ranks defining second NMF compartment | |
| specificity | the fold-change the ranks of interest must have relative to remaining ranks to be included | 2 |
| threshold | fractional threshold (relative to max) to consider a prey enriched in a compartment | 0.5 |

## Example file formats

### basisMatrix
```
variable,1,2,3
AAAS,0.0,0.18373784550412228,0.0
AAK1,0.13963834317658874,0.0,0.0
AAR2,0.0,0.016517285371126216,0.034697999230769466
AARS2,0.0,0.0,0.001598769548285137
```

### Output
* heat map as svg
* legend for fold enrichment