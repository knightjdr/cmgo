# Module name: `analysis-dbgenes`

> Generates a txt file of official genes used for SAINT

| parameters | definition | default |
|------------|------------|---------|
| ncbigene | NCBI gene file with gene names  | |
| outFile | output list of gene names | db-genes.txt |

## Example file formats

Download NCBI gene information from: `ftp://ftp.ncbi.nlm.nih.gov/gene/DATA/gene_info.gz`

### gene_info
```
#tax_id	GeneID	Symbol	LocusTag	Synonyms	dbXrefs	chromosome	map_location	description	type_of_gene	Symbol_from_nomenclature_authority	Full_name_from_nomenclature_authority	Nomenclature_status	Other_designations	Modification_date	Feature_type
192	36107229	AMK58_RS00870	AMK58_RS00870	AMK58_00870	-	-	-	2-polyprenylphenol 6-hydroxylase	protein-coding	-	-	-	-	20180309	-
9606	1	A1BG	-	A1B|ABG|GAB|HYST2477	MIM:138670|HGNC:HGNC:5|Ensembl:ENSG00000121410	19	19q13.43	alpha-1-B glycoprotein	protein-coding	A1BG	alpha-1-B glycoprotein	O	alpha-1B-glycoprotein|HEL-S-163pA|epididymis secretory sperm binding protein Li 163pA	20190617	-
9606	2	A2M	-	A2MD|CPAMD5|FWP007|S863-7	MIM:103950|HGNC:HGNC:7|Ensembl:ENSG00000175899	12	12p13.31	alpha-2-macroglobulin	protein-coding	A2M	alpha-2-macroglobulin	O	alpha-2-macroglobulin|C3 and PZP-like alpha-2-macroglobulin domain-containing protein 5|alpha-2-M	20190817	-
```

### Output
* txt files with one gene names per line