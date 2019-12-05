# Module name: `interaction-turnoverbyrank`

> Calculates the average turnover rate for preys, ranked by
> their control-subtracted length-adjusted spectral counts

| parameters | definition | default |
|------------|------------|---------|
| fdr | FDR threshold | 0.01 |
| turnoverFile | file with turnover rates per gene | |
| outFile | output file | turnover-by-rank.txt |
| saint | SAINT file | |

## Example file formats

### turnoverFile

PMID:     29414762, Supplementary table S3

### Output
* ` turnover-by-rank.txt`: table with turnover rate (mean and SD) and the number of genes
with a turnover rate at each rank

```
prey rank turnover rate (mean) turnover rate (SD)  genes with turnover data
```