# Module name: `organelle-isolation`

> Calculates how isolated an NMF compartment is, i.e. for genes localizing to the compartment
> the number of edges between members of the compartment relative to the number of edges to proteins
> outside the compartment.

| parameters | definition | default |
|------------|------------|---------|
| abundanceCap | value cap on absolute heatmap | 1000 |
| basisMatrix | NMF basis matrix | |
| correlationCutoff | cutoff to consider an interaction pair | 0.9 |
| nmfLocalization | Text file with assigned NMF rank and score for each gene | |
| nmfSummary | NMF rank summary file | |
| outFile | output file name | organelle-isolation.txt |
| svgFile | output file name for heatmap | organelle-isolation.svg |

## Example file formats

### basisMatrix
```
variable,1,2,3
AAAS,0.0,0.18373784550412228,0.0
AAK1,0.13963834317658874,0.0,0.0
AAR2,0.0,0.016517285371126216,0.034697999230769466
AARS2,0.0,0.0,0.001598769548285137
```

### nmfLocalization
```
gene	rank	score
AAAS	13	0.3326865	
AAK1	5	0.1396383	
AAR2	19	0.034698	
AARS2	6	0.2251458	
AASDH	2	0.09391108
```

### nmfSummary
```
rank	term	displayname	go	synonyms	ic
1	[cell junction]	[cell junction]	[GO:0030054]	[]	[1.166]
2	[chromosome]	[chromatin]	[GO:0005694]	"[[chromatid, interphase chromosome, prophase chromosome]]"	[1.256]
```

## Output

`organelle-isolation.txt`: for each compartment lists the number of edges within the compartment and outside
`heatmap`: heatmap showing the absolute or relative number of edges shared between each compartments