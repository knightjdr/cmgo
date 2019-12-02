# Module name: `assessment-countgo`

> Count the number of genes with a GO term specified by the supplied namespace.

| parameters | definition | default |
|------------|------------|---------|
| genes | list of genes with assigned localizations | |
| goAnnotations | GO annotations in .gaf format | |
| namespace | GO namespace to use, one of BP, CC, MF | CC |

## Example file formats

### goAnnotations
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

### genes
File with a header and genes names in the first column.
