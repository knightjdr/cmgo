# Module name: `assessment-localization`

> Report on which localizations are previously known.

| parameters | definition | default |
|------------|------------|---------|
| goAnnotations | GO annotations in .gaf format | |
| goHierarchy | GO hierarchy in .obo format | |
| localization | list of genes with assigned localizations | |
| namespace | GO namespace to use, one of BP, CC, MF | CC |
| outFile | output file name | localization-known.txt |
| outFileSummary | output file name for summary statistics | summary.txt |

## Example file formats

### annotationsFile
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

### localization
```
gene	term(s)	ID(s)	precision
AAAS	nuclear outer membrane-endoplasmic reticulum membrane network	GO:0042175	0.69
AAK1	cytosol	GO:0005829	0.67
AAR2	nucleoplasm	GO:0005654	0.66
AARS2	mitochondrial matrix	GO:0005759	0.78
AASDH	nucleoplasm	GO:0005654	0.79
AASS	mitochondrial matrix	GO:0005759	0.76
```

### Output
* `localization-known.txt`: the input file with a boolean indicate if known
* `summary.txt`: summary statistics