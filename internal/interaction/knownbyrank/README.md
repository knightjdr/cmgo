# Module name: `interaction-knownbyrank`

> Calculates the proportion of known interactors for the Nth best prey across
> all baits in a SAINT report. Prey spectral counts are control subtracted
> and length normalized to determine their rank.

| parameters | definition | default |
|------------|------------|---------|
| biogrid | file with BioGRID interactions | |
| fdr | FDR threshold | 0.01 |
| intact | file with IntAct interactions | |
| outFile | output file | known-by-rank.txt |
| saint | SAINT file | |
| species | taxonID for IntAct entries (should be passed as string) | 9606 |

## Example file formats

### biogrid
```
#BioGRID Interaction ID	Entrez Gene Interactor A	Entrez Gene Interactor B	BioGRID ID Interactor A	BioGRID ID Interactor B	Systematic Name Interactor A	Systematic Name Interactor B	Official Symbol Interactor A	Official Symbol Interactor B
103	6416	2318	112315	108607	-	-	MAP2K4	FLNC
117	84665	88	124185	106603	-	-	MYPN	ACTN2
183	90	2339	106605	108625	-	-	ACVR1	FNTA
278	2624	5371	108894	111384	-	-	GATA2	PML
418	6118	6774	112038	112651	RP4-547C9.3	-	RPA2	STAT3
586	375	23163	106870	116775	-	-	ARF1	GGA3
612	377	23647	106872	117174	-	-	ARF3	ARFIP2
617	377	27236	106872	118084	-	-	ARF3	ARFIP1
663	54464	226	119970	106728	-	-	XRN1	ALDOA
```

### intact
```
#ID(s) interactor A     ID(s) interactor B      Alt. ID(s) interactor A Alt. ID(s) interactor B Alias(es) interactor A  Alias(es) interactor B  Interaction detection method(s) Publication 1st author(s)       Publication Identifier(s)       Taxid interactor A      Taxid interactor B      Interaction type(s)     Source database(s)      Interaction identifier(s)       Confidence value(s)     Expansion method(s)     Biological role(s) interactor A Biological role(s) interactor B Experimental role(s) interactor A       Experimental role(s) interactor B       Type(s) interactor A    Type(s) interactor B    Xref(s) interactor A    Xref(s) interactor B    Interaction Xref(s)     Annotation(s) interactor A      Annotation(s) interactor B      Interaction annotation(s)       Host organism(s)        Interaction parameter(s)        Creation date   Update date     Checksum(s) interactor A        Checksum(s) interactor B        Interaction Checksum(s) Negative        Feature(s) interactor A Feature(s) interactor B Stoichiometry(s) interactor A   Stoichiometry(s) interactor B   Identification method participant A     Identification method participant B
uniprotkb:P49418        uniprotkb:O43426        intact:EBI-7121510|uniprotkb:Q75MK5|uniprotkb:Q75MM3|uniprotkb:A4D1X9|intact:MINT-109264|uniprotkb:O43538|uniprotkb:A4D1X8|uniprotkb:Q75MJ8|uniprotkb:Q8N4G0    intact:EBI-2821539|uniprotkb:O43425|uniprotkb:O94984|uniprotkb:Q4KMR1   psi-mi:amph_human(display_long)|uniprotkb:AMPH(gene name)|psi-mi:AMPH(display_short)|uniprotkb:AMPH1(gene name synonym)
```

### Output
* `known-by-rank.txt`: table with the proportion for each Nth best prey that are previously
known interactors.

In the below example, there are 10 baits with at least 1 prey and 80% of the first preys are
previously known interactors. There are 7 baits with at least 2 preys and 70% of the second
best preys are previously known.

```
prey rank proportion number of baits  known pairs
1 0.8 10  8 "geneA-geneB, geneC-GeneD"
2 0.7 10  7 "geneA-geneB, geneC-GeneD"
```