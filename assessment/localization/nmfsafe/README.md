# Module name: `nmf-safe`

> Check concordance between NMF and SAFE localizations.

| parameters | definition | default |
|------------|------------|---------|
| annotationsFile | GO annotations in .gaf format | |
| goHierarchy | GO hierarchy in .obo format | |
| nmfLocalization | Text file with assigned NMF rank and score for each gene | |
| nmfSummary | NMF rank summary file | |
| outFile | output file name | concordance.txt |
| safeLocalization | Node properties file output from SAFE | |
| safeSummary | SAFE domain summary file | |
| summaryFile | file with summary statistics | summary.txt

## Example file formats

## annotationsFile
```
!Documentation about this header can be found here: https://github.com/geneontology/go-site/blob/master/docs/gaf_validation.md
!
UniProtKB	A0A024R161	DNAJC25-GNG10		GO:0003924	GO_REF:0000002	IEA	InterPro:IPR001770	F	Guanine nucleotide-binding protein subunit gamma	DNAJC25-GNG10|hCG_1994888	protein	taxon:9606	20190112	InterPro		
UniProtKB	A0A024R161	DNAJC25-GNG10		GO:0007186	GO_REF:0000002	IEA	InterPro:IPR001770|InterPro:IPR015898|InterPro:IPR036284	P	Guanine nucleotide-binding protein subunit gamma	DNAJC25-GNG10|hCG_1994888	protein	taxon:9606	20190112	InterPro
UniProtKB	Q9BUL8	PDCD10		GO:0000139	GO_REF:0000039	IEA	UniProtKB-SubCell:SL-0134	C	Programmed cell death protein 10	PDCD10|CCM3|TFAR15	protein	taxon:9606	20190112	UniProt		
UniProtKB	Q9BUL8	PDCD10		GO:0001525	GO_REF:0000037	IEA	UniProtKB-KW:KW-0037	P	Programmed cell death protein 10	PDCD10|CCM3|TFAR15	protein	taxon:9606	20190112	UniProt		
UniProtKB	Q9BUL8	PDCD10		GO:0005515	PMID:16189514	IPI	UniProtKB:O00506	F	Programmed cell death protein 10	PDCD10|CCM3|TFAR15	protein	taxon:9606	20190113	IntAct		
UniProtKB	Q9BUL8	PDCD10		GO:0005515	PMID:16189514	IPI	UniProtKB:Q9Y6E0	F	Programmed cell death protein 10	PDCD10|CCM3|TFAR15	protein	taxon:9606	20190113	IntAct
UniProtKB	Q9BUL9	RPP25		GO:0001682	Reactome:R-HSA-5696810	TAS		P	Ribonuclease P protein subunit p25	RPP25	protein	taxon:9606	20151024	Reactome		
UniProtKB	Q9BUL9	RPP25		GO:0003723	PMID:22658674	HDA		F	Ribonuclease P protein subunit p25	RPP25	protein	taxon:9606	20140203	UniProt		
UniProtKB	Q9BUL9	RPP25		GO:0005515	PMID:15096576	IPI	UniProtKB:O95707	F	Ribonuclease P protein subunit p25	RPP25	protein	taxon:9606	20190113	IntAct		
UniProtKB	Q9BUL9	RPP25		GO:0005515	PMID:21044950	IPI	UniProtKB:Q9NYB0	F	Ribonuclease P protein subunit p25	RPP25	protein	taxon:9606	20190113	IntAct		
UniProtKB	Q9BUL9	RPP25		GO:0005654	GO_REF:0000052	IDA		C	Ribonuclease P protein subunit p25	RPP25	protein	taxon:9606	20141106	HPA	
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

### nmfLocalization
```
gene	rank	score
AAAS	13	0.3326865	
AAK1	5	0.1396383	
AAR2	19	0.034698	
AARS2	6	0.2251458	
AASDH	2	0.09391108
```

### nmfSummary and safeSummary
```
rank	term	displayname	go	synonyms	ic
1	[cell junction]	[cell junction]	[GO:0030054]	[]	[1.166]
2	[chromosome]	[chromatin]	[GO:0005694]	"[[chromatid, interphase chromosome, prophase chromosome]]"	[1.256]
```

### safeLocalization
```
## 
## This file lists the properties of all nodes in the network.
## 

Node label	Node label ORF	Domain (predominant)	Neighborhood score [max=1, min=0] (predominant)	Total number of enriched domains	Number of enriched attributes per domain
VAMP3	VAMP3	20	0.600	1	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,0,0,0
SNAP29	SNAP29	1	0.263	0	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
CDCA3	CDCA3	20	1.000	1	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,61,0,0,0,0
```

### Output
* txt file: for each gene list NMF and SAFE term, is it known, and whether they agree
* txt file with summary statistics