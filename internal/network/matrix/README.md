# Module name: `network-matrix`

> Generate network from a matrix.

| parameters | definition | default |
|------------|------------|---------|
| colorList | text file with hex colors | |
| cutoff | cutoff for matrix value that should be considered an edge | 0.01 |
| localizations | ordered list of all possible localizations | |
| matrix | matrix with genes as rows and localizations as columns | |
| nodeLocalizations | list of primary localizations for genes | |
| outFile | output file with all node pairs and edge weight (correlation) | matrix.txt |
| outFileNetwork | output file name Cytoscape network | matrix.cyjs |

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

### matrix
```
gene	GO:0005694	GO:0016604	GO:0005635
AAAS	0.0000	0.0000	0.2500
AAK1	0.0000	0.2200	0.6700
```

### Output
* `matrix.cyjs`: Cytoscape formatted file with nodes placed on a grid
* `matrix.txt`: node pairs
