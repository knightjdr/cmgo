# Module name: `enrichment-heatmap`

> Creates a heat map visualizing region (domain/motif) enrichment across NMF or SAFE compartments.

| parameters | definition | default |
|------------|------------|---------|
| compartmentSummary | files with names defining compartments | |
| enrichmentFile | text file with compartment name and enriched domains/motifs | |
| outFile | output file name | region-heatmap.svg |
| pValue | pValue cutoff for including a region | 0.01 |

## Example file formats

### compartmentSummary
```
rank	term	displayname	go	synonyms	ic
1	[cell junction]	[cell junction]	[GO:0030054]	[]	[1.166]
2	[chromosome]	[chromatin]	[GO:0005694]	"[[chromatid, interphase chromosome, prophase chromosome]]"	[1.256]
```

### enrichmentFile
```
rank	term	matched	background_size	fold enrichment	pvalue	adj. pvalue	bhfdr	genes
1	PDZ	23	48	8.906	1.62980740014973e-13	4.27009538839229e-11	3.81679389312977e-05	AFDN,CASK,DLG1
1	FERM_C	12	17	13.120	2.1798629399977e-12	2.85562045139698e-10	7.63358778625954e-05	EPB41,EPB41L1,EPB41L2
1	FERM_N	12	18	12.391	6.23068998013292e-12	5.44146924931608e-10	0.000114503816793893	EPB41,EPB41L1,EPB41L2
2	KRAB	24	28	9.905	2.32442462867683e-22	1.02972011050383e-19	2.25733634311512e-05	POGK,RBAK,ZFP1
2	zf-C2H2	30	65	5.333	2.99760216648792e-15	6.63968879877075e-13	4.51467268623025e-05	CTCF,MAZ,PRDM15
2	Bromodomain	13	24	6.259	1.29285132936678e-08	1.70043585864132e-06	6.77200902934537e-05	ATAD2,BPTF,BRD1
2	PWWP	10	14	8.254	1.53538226513889e-08	1.70043585864132e-06	9.0293453724605e-05	BRD1,BRPF3,DNMT3A,GLYR1
```
