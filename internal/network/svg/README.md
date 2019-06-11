# Module name: `network-svg`

> Create an svg network for NMF or SAFE

| parameters | definition | default |
|------------|------------|---------|
| colorList | text file with hex colors | |
| localizations | node localizations | |
| nodeCoordinates | text file with list of nodes and x, y coordinates |
| outFile | output file name | network.svg |

## Example file formats

### colorList
```
#C0B9B2
#1CE6FF
#FF4A46
#008941
```

### localizations

#### NMF
```
gene	rank	score
AAAS	13	0.3326865	
AAK1	5	0.1396383	
AAR2	19	0.034698	
AARS2	6	0.2251458	
AASDH	2	0.09391108	
AASS	4	0.3192908
```

#### SAFE
```
## 
## This file lists the properties of all nodes in the network.
## 

Node label	Node label ORF	Domain (predominant)	Neighborhood score [max=1, min=0] (predominant)	Total number of enriched domains	Number of enriched attributes per domain
VAMP3	VAMP3	20	0.600	1	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,0,0,0
SNAP29	SNAP29	1	0.263	0	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
CDCA3	CDCA3	20	1.000	1	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,61,0,0,0,0
ZC3HAV1	ZC3HAV1	19	1.000	1	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,17,0,0,0,0,0
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
* 