# Module name: `network-tsnecytoscape`

> Convert a tSNE coordinate system into a .cyjs file for Cytoscape

| parameters | definition | default |
|------------|------------|---------|
| colorList | text file with hex colors | |
| localizations | ordered list of all possible localizations | |
| nodeCoordinates | text file with list of nodes and x, y coordinates |
| nodeLocalizations | list of primary localizations for genes | |
| outFile | output file name | map.cyjs |
| width | the ideal width for the map | 1000 |

## Example file formats

### colorList
```
#C0B9B2
#1CE6FF
#FF4A46
#008941
```

### nodeCoordinates
```
gene	x	y
AAAS	68.5533816047702	-33.9816908466039
AAK1	-28.8672128420763	-43.6402491976645
AAR2	2.2808526181511	-6.97152015675554
AARS2	89.5765051986436	67.0643311315994
AASDH	8.41236940140679	35.7167588576447
AASS	-70.8218291970252	64.4050187741019
AATF	-15.8239071660148	51.4522241800074
ABCA3	-19.8932838810909	-3.05943018550193
ABCB1	-16.1846453937935	91.5103672699896
ABCB10	-52.4758726133164	73.6495265978388
```

### Output
* `map.cyjs`: map in .cyjs format for Cytoscape