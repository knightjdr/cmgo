# Module name: `nmf-moonlighting`

> Calculates a moonlighting score for preys in an NMF basis matrix

| parameters | definition | default |
|------------|------------|---------|
| dissimilarityFile | a binary upper triangular matrix of rank dissimilarity | |
| minimumNmfScore | minimum NMF score for a primary rank to include a prey in the analysis | 0.15 |
| nmfBasis | NMF basis matrix | |
| outFile | output file | moonlighting.txt |

## Example file formats

### nmfBasis
```
variable,1,2,3
AAAS,0.0,0.18373784550412228,0.0
AAK1,0.13963834317658874,0.0,0.0
AAR2,0.0,0.016517285371126216,0.034697999230769466
AARS2,0.0,0.0,0.001598769548285137
```

### dissimilarityFile
```
  1 2 3
1 0 1 1
2   0 1
3     0
```

### Output
* `moonlighting.txt`: table with prey, moonlighting score, primary rank, secondary and tertiary