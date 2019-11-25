# Module name: `nmf-moonlighting`

> Calculates a moonlighting score for preys in an NMF basis matrix.

The primary localization must have a score >= minRankValue to be included in the analysis,
and the secondary localization must also be >= minRankValue to be a valid secondary localization.

| parameters | definition | default |
|------------|------------|---------|
| basisMatrix | NMF basis matrix | |
| dissimilarityMatrix | a binary upper triangular matrix of rank dissimilarity | |
| minRankValue | minimum NMF score for a primary rank to include a prey in the analysis | 0.15 |
| nmfSummary | NMF rank summary file | |
| outFileHeatmap | heat map of primary v secondary rank moonlighting | heatmap.svg |
| outFileMatrix | raw values from heatmap | matrix.txt |
| outFileScores | output file | moonlighting.txt |

## Example file formats

### basisMatrix
```
variable,1,2,3
AAAS,0.0,0.18373784550412228,0.0
AAK1,0.13963834317658874,0.0,0.0
AAR2,0.0,0.016517285371126216,0.034697999230769466
AARS2,0.0,0.0,0.001598769548285137
```

### dissimilarityMatrix
```
  1 2 3
1 0 1 1
2   0 1
3     0
```

### nmfSummary
```
rank	term	displayname	go	synonyms	ic
1	[cell junction]	[cell junction]	[GO:0030054]	[]	[1.166]
2	[chromosome]	[chromatin]	[GO:0005694]	"[[chromatid, interphase chromosome, prophase chromosome]]"	[1.256]
```

### Output
* `heatmap.svg`: heat map for the `matrix.txt` file
* `matrix.txt`: matrix with the number of times a rank scoring as a primary localization (row) was
detected with a rank as a secondary localization (column)
* `moonlighting.txt`: table with prey, moonlighting score, primary rank and secondary rank