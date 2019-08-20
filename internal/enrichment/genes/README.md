# Module name: `enrichment-genes`

> Perform a g:Profiler enrichment on a list of genes

| parameters | definition | default |
|------------|------------|---------|
| background | list of background genes | |
| genes | list of gene names  | |
| namespace | GO namespace to check for enrichment | CC |
| outFile | output of enrichment | enrichment.txt |

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

### Output
* txt file with g:Profiler enrichment