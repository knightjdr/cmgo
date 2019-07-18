# Module name: `summary-notsignificant`

> Generates a list of preys that were not significant with any bait in a SAINT file.

| parameters | definition | default |
|------------|------------|---------|
| fdr | fdr threshold for significance (inclusive) | 0.01 |
| outFile | output file name | not-significant.txt |
| saint | saint file name | 0 |

## Example file formats

### Output
* txt file list baits a preys was seen with, its maximum spectral count (from Spec column), the best FDR it had and its average (mean) value in controls calculated for the row where the maximum spectral count was found