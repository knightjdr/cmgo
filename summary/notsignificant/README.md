# Module name: `summary-notsignificant`

> Subset an NMF basis (prey) matrix to only include preys enriched in specified compartments.

| parameters | definition | default |
|------------|------------|---------|
| fdr | fdr threshold for significance (inclusive) | 0.01 |
| outFile | output file name | not-significant.txt |
| saint | saint file name | 0 |

## Example file formats

### Output
* txt file list baits a preys was seen with, its maximum spectral count (from Spec column), the best FDR it had and its average (mean) value in controls calculated for the row where the maximum spectral count was found