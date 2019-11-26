# Module name: `assessment-transmembrane`

> Get information about the transmembrane proteins residing between two NMF
> compartments.

| parameters | definition | default |
|------------|------------|---------|
| basisMatrix | NMF basis matrix | |
| cytosolicBaits | baits that were used to label the cytosolic compartment (comma-separated list) | |
| cytosolicCompartments | ranks defining cytosolic compartments (comma-separated list) | |
| fdr | FDR | 0.01 |
| minRankValue | minimum NMF score for a primary rank to include a prey in the analysis | 0.15 |
| lumenalBaits | baits that were used to label the lumenal compartment (comma-separated list) | |
| lumenalCompartments | ranks defining lumenal compartments (comma-separated list) | |
| outFile | list of preys with transmembrane domains and associated information | transmembrane.txt |
| saint | SAINT file | |

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
* `transmembrane.txt`