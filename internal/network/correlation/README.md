# Module name: `network-correlation`

> Calculate correlation between matrix rows and output Cytoscape formatted file for viewing network.

| parameters | definition | default |
|------------|------------|---------|
| colorList | text file with hex colors | |
| cutoff | correlation cutoff to define edges | 0 |
| edgesPerNode | ideal edge number between nodes when not setting a cutoff | 20 |
| localizations | ordered list of all possible localizations | |
| maxEdges | maximum number of edges to allow per node when specifying a cutoff | 0 |
| nodeLocalizations | list of primary localizations for genes | |
| nodeProfiles | localization profiles (matrix) | |
| outFile | output file with all node pairs and edge weight (correlation) | corr.txt |
| outFileNetwork | output file name Cytoscape network | corr.cyjs |

If `cutoff` is > 0, it will be used to define edges, otherwise a cutoff will be sought
to achieve `edgesPerNode`.

## Example file formats

### colorList
```
#C0B9B2
#1CE6FF
#FF4A46
#008941
```

### localization
```
GO:0005654	nucleoplasm
GO:0005694	chromosome
GO:0005730	nucleolus
GO:0005741	mitochondrial outer membrane
GO:0005743	mitochondrial inner membrane
```

### nodeLocalizations
```
gene	GO:0005694	GO:0016604	GO:0005635
AAAS	0.0000	0.0000	0.2500
AAK1	0.0000	0.2200	0.6700
```

### Output
* `corr.cyjs`: Cytoscape formatted file with nodes placed on a grid
* `corr.txt`: node pairs
