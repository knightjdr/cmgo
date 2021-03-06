# Module name: `assessment-prediction`

> Assign a prediction score for every prey localized by NMF or SAFE

| parameters | definition | default |
|------------|------------|---------|
| baitExpected | expected bait localizations | |
| compartmentsText | text-mining score from Comportments DB | |
| domainsPerCompartment | file with a list of domains for each NMR rank or SAFE domain | |
| domainsPerGene | file with a list of domains for each gene | |
| fdr | FDR for significant prey in SAINT file | |
| fractionation | predictions from fractionation studies of Christoforou and Itzhak | |
| goHierarchy | GO hierarchy in .obo format | |
| hpa | human protein atlas predictions | |
| outFile | output file | prediction-score.txt |
| predictions | file with localization predictions by NMF rank or SAFE domain | |
| predictionSummary | summary information for NMF or SAFE ranks | |
| predictionType | type of prediction: NMF or SAFE | nmf |
| saint | SAINT file | |

| uniprot | UniProt file in .dat format | |

## Example file formats

### baitExpected
```
id	bait	term	GO_ID
1	AARS2	mitochondrial matrix	GO:0005759
2	ACBD5	peroxisome	GO:0005777
3	ACTB	actin cytoskeleton	GO:0015629
```

### compartmentsText
```
ENSP00000255087	MTL5	GO:0042645	Mitochondrial nucleoid	1.010	0.505	
ENSP00000255087	MTL5	GO:0000229	Cytoplasmic chromosome	1.009	0.504	
ENSP00000255087	MTL5	GO:0045495	Pole plasm	1.004	0.502	
ENSP00000255108	DPH2	GO:0005840	Ribosome	3.332	1.666	
ENSP00000255108	DPH2	GO:1990904	Ribonucleoprotein complex	3.077	1.538
```

### domainsPerCompartment
```
rank	term	matched	background_size	fold enrichment	pvalue	adj. pvalue	bhfdr	genes
1	PDZ	23	48	8.906	1.62980740014973e-13	4.27009538839229e-11	3.81679389312977e-05	AFDN,CASK
1	FERM_C	12	17	13.120	2.1798629399977e-12	2.85562045139698e-10	7.63358778625954e-05	EPB41,EPB41L5
```

### domainsPerGene
```
#<seq id> <alignment start> <alignment end> <envelope start> <envelope end> <hmm acc> <hmm name> <type> <hmm start> <hmm end> <hmm length> <bit score> <E-value> <clan>
O95863	153	176	153	177	PF13912	zf-C2H2_6	PfamLive::Result::SequenceOntology=HASH(0x108198b8)	1	24	27	20.50	0.9	CL0361
O95863	208	230	208	230	PF00096	zf-C2H2	PfamLive::Result::SequenceOntology=HASH(0x89b8a80)	1	23	23	26.70	0.013	CL0361
O95863	181	202	180	202	PF00096	zf-C2H2	PfamLive::Result::SequenceOntology=HASH(0x89b8a80)	2	23	23	23.40	0.14	CL0361
```

### fractionation
```
gene	localization Christoforou	localization Itzhak	ID Christoforou	ID Itzhak	combined
AARS	Cytosol	Mostly Cytosolic	GO:0005829	GO:0005829	GO:0005829
AARS2	Mitochondrion	Mitochondrion Matrix	GO:0005739	GO:0005759	GO:0005739;GO:0005759
AARSD1	Cytosol	Mostly Cytosolic	GO:0005829	GO:0005829	GO:0005829
```

### goHierarchy
```
[Term]
id: GO:0005652
name: nuclear lamina
namespace: cellular_component
def: "The fibrous, electron-dense layer lying on the nucleoplasmic side of the inner membrane of a cell nucleus." [ISBN:0198506732, ISBN:0716731363]
xref: NIF_Subcellular:sao1455996588
xref: Wikipedia:Nuclear_lamina
is_a: GO:0044428 ! nuclear part
relationship: part_of GO:0034399 ! nuclear periphery

[Term]
id: GO:0005654
name: nucleoplasm
namespace: cellular_component
def: "That part of the nuclear content other than the chromosomes or the nucleolus." [GOC:ma, ISBN:0124325653]
subset: goslim_chembl
subset: goslim_generic
subset: goslim_plant
xref: NIF_Subcellular:sao661522542
xref: Wikipedia:Nucleoplasm
is_a: GO:0044428 ! nuclear part
relationship: part_of GO:0031981 ! nuclear lumen
```

### hpa
```
gene	terms	go_id
TSPAN6	Cytosol	GO:0005829
SCYL3	Microtubules;Nuclear bodies	GO:0015630;GO:0016604
C1orf112	Mitochondria	GO:0005739
FGR	Aggresome;Plasma membrane	GO:0016235;GO:0005886
CFH	Vesicles	GO:0043231
```

### predictions

#### NMF
```
gene	rank	score
AAAS	13	0.3326865	
AAK1	5	0.1396383	
AAR2	19	0.034698	
AARS2	6	0.2251458	
AASDH	2	0.09391108
```

#### SAFE
```
Node label	Node label ORF	Domain (predominant)	Neighborhood score [max=1, min=0] (predominant)	Total number of enriched domains	Number of enriched attributes per domain
VAMP3	VAMP3	20	0.600	1	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,0,0,0
SNAP29	SNAP29	1	0.263	0	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
CDCA3	CDCA3	20	1.000	1	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,61,0,0,0,0
```

### predictionSummary
```
rank	term	displayname	go	synonyms	ic
1	[cell junction]	[cell junction]	[GO:0030054]	[]	[1.166]
2	[chromosome]	[chromatin]	[GO:0005694]	"[[chromatid, interphase chromosome, prophase chromosome]]"	[1.256]
```

## Output
`prediction-score.txt`: output prediction score components for each prey