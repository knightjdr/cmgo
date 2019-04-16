# Module name: `organelle-shareddomain`

Takes two lists of proteins and outputs metrics on the domains/motifs shared between common preys.

| parameters | definition | default |
|------------|------------|---------|
| compartmentFile | JSON file with lists of proteins to compare | |
| outFile | output file name | organelle-shared.txt |
| regionFile | text file with protein name and domains/motifs | |
| saintFile | SAINT file | |

## Example file formats

### compartmentFile
```
[
  {
    "name": "compartment 1",
    "proteins": ["a", "b" ,"c"]
  },
  {
    "name": "compartment 2",
    "proteins": ["d", "e" ,"f"]
  }
]
```

### regionFile (tab-separated, no header)
```
A1BG	disorder
A1BG	sig_p
A1CF	disorder
A1CF	low_complexity
A2M	disorder
A2M	low_complexity
A2M	sig_p
A2ML1	disorder
A2ML1	low_complexity
A2ML1	sig_p
A3GALT2	low_complexity
A3GALT2	transmembrane
```

### Output (example)
```
```