# Module name: `organelle-overlap`

> Takes two lists of proteins and a txt file with similarity scores between proteins and outputs metrics on the similarity within and between the lists.

| parameters | definition | default |
|------------|------------|---------|
| compartmentFile | JSON file with lists of proteins to compare | |
| outFile | output file name | organelle-overlap.txt |
| similarityFile | three-column tsv file (with header) | |

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

### similarityFile (tab-separated)
```
source  target  similarity
a   a   1
a   b   0.5
a   c   0.25
b   a   0.5
b   b   1
b   c   0.35
c   a   0.25
c   b   0.35
c   c   1
```

### Output (example)
```
median  mean    min     max
ER lumen        0.541   0.448   0.087   0.699
ER membrane     0.429   0.380   0.035   0.706
between 0.069   0.116   0.013   0.655
```