# Module name: `lba-localize`

> Localize preys based on enriched terms

| parameters | definition | default |
|------------|------------|---------|
| enrichment | list of enriched terms for each gene | |
| localization | list of GO IDs that are considered valid localizations | |
| outFilePrimary | output file name for prey localization data | lba-primary.txt |

## Example file formats

### enrichment

```
symbol	Entrez	Refseq	UniProt	GO ID	GO term	p-value	recall	precision	query size	term size
RBM15	64783	NP_001188474.1	Q96T37	GO:0044428	nuclear part	4.243467e-19	0.01	0.88	50	4521
RBM15	64783	NP_001188474.1	Q96T37	GO:0031981	nuclear lumen	4.508720e-18	0.01	0.84	50	4142
RBM15	64783	NP_001188474.1	Q96T37	GO:0070013	intracellular organelle lumen	6.906764e-14	0.01	0.84	50	5283
RBM15	64783	NP_001188474.1	Q96T37	GO:0043233	organelle lumen	6.906764e-14	0.01	0.84	50	5283
CSNK1G3	1456	NP_001026982.1	Q9Y6M4	GO:0005886	plasma membrane	1.838539e-19	0.01	0.94	50	5539
CSNK1G3	1456	NP_001026982.1	Q9Y6M4	GO:0071944	cell periphery	5.044889e-19	0.01	0.94	50	5662
```

### localization
```
GO:0005654	nucleoplasm
GO:0005694	chromosome
GO:0005730	nucleolus
GO:0005741	mitochondrial outer membrane
GO:0005743	mitochondrial inner membrane
```

### Output
* `lba-primary.txt`: a file with the GO ID(s) and term name(s) for every prey's assigned localization
* `lba-profile.txt`: matrix showing the precision for each gene across specified localizations. Localizations are
in the same order as they are input