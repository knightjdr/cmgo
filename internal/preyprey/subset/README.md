# Module name: `preyprey-subset`

> Subset a ProHits-viz interactive file to grab a cluster

| parameters | definition | default |
|------------|------------|---------|
| genes | list of gene names that identify the cluster  | |
| heatmap | .tsv file from ProHits-viz  | |
| outFile | output list of gene names | cluster.tsv |

## Example file formats

### genes

```
WDR3
NOP56
DKC1
NOLC1
TCOF1
URB1
DHX33
TRMT1L
```

### heatmap

```
row	column	value	params
SNAP29	SNAP29	1	{"type": "heatmap", "kind": "Prey vs prey", "xAxis": "Prey", "yAxis": "Prey", "filterType": 0, "abundance": "Correlation - Spec"}
SNAP29	PALD1	0.571787038973909	
SNAP29	APBB1	0.621831039504651	
SNAP29	EHBP1	0.619991568749013	
SNAP29	KCNB2	0.392832468690144	
SNAP29	UBA52	0.557246514867496	
SNAP29	VANGL2	0.59688567365931	
```

### Output
* tsv file for input to ProHits-viz viewer