# Module name: `interaction-rankaverage`

> Calculates the average bait interaction rank for a list of preys.

| parameters | definition | default |
|------------|------------|---------|
| fdr | FDR threshold | 0.01 |
| outFile | output file | prey-rank-average.txt |
| preys | list of preys | |
| saint | SAINT file | |

## Example file formats

### preys
Must have a header, with first column having gene names.
```
gene
preyA
preyB
preyC
```

### Output
* ` prey-rank-average.txt`
```
prey  mean  sd  interaction ranks
preyA 3.5 0.707 3,4
```

preyA interacted with two baits, and was the 3 and 4th best prey for those baits
after control-subtracting and length adjusting spectral counts