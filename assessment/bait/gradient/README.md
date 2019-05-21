# Module name: `bait-gradient`

> Takes an ordered bait list and expected localizations for the baits and outputs a gradient showing
> adjacent bait similarity: black if the two adjacent baits share an expected localization, grey
> if only one does and white otherwise.

| parameters | definition | default |
|------------|------------|---------|
| baitList | list of baits in order of similarity | |
| expectedLocalizations | .txt file with bait name and expected localizations  | |
| outFile | output file name | bait-gradient.svg |

## Example file formats

## baitList
```
HNRNPA1
RPS20
SERBP1
FBL
RPS6
RPL31	
```

### expectedLocalizations
```
id	bait	localization
1	AARS2	mitochondrial matrix
2	ACBD5	peroxisome
3	ACTB	actin cytoskeleton
12	ANAPC2	"cytoplasm;nucleoplasm;nucleus"
13	ANK3	"cell junction;plasma membrane"
```

### Output
* svg gradient in same order as input bait list
