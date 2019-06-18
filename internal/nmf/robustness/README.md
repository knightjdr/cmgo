# Module name: `nmf-robustness`

> Evaluates how robust each rank's GO terms are based on the genes used
> to define it. It will first define the rank by performing a GO enrichment
> using at most maxGenesPerRank, and then repeat this when using 90%, 80%...
> of the maxGenesPerRank. For each evaluation it will calculate the RBD between
> the list generated for maxGenesPerRank and the fractional list.

| parameters | definition | default |
|------------|------------|---------|
| basisMatrix | NMF basis matrix | |
| maxGenesPerRank | maximum number of genes to use per rank for enrichments | 100 |
| minRankValue | a prey must have an NMF value at or above this value to be used for enrichment | 0.25 |
| outFile | output file name for RBD data points | robustness.txt |
| outFileSummary | output file name for summary (mean, SD) statistics | summary.txt |
| percentiles | comma-separated string defining percentiles to test, e.g "0.9,0.8" | |
| persistence | RBO persistence value | 0.9 |
| replicates | number of replicates to perform for each percentile | 3 |
| withinRankMax | if a prey has an NMF value within this % of max in its non-primary rank, it can be used for enrichment | 0.75 |

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
* `robustness.txt`: RBD data points for each rank, percentile and replicate
* `summary.txt`: mean and SD for each rank and percentile