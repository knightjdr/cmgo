# Module name: `control-preys`

> Output statistics on prey proteins in controls using SAINT input files.

| parameters | definition | default |
|------------|------------|---------|
| bait | bait.dat file | |
| enrichmentLimit | number of preys to user for GO enrichment after sorting | 200 |
| inter | bait.dat file | |
| outFile | output file | control-preys.txt |
| outFileEnrichment | output file for GO enrichment | enrichment.txt |
| prey | bait.dat file | |

## Example file formats

### bait

`bait.dat` file from SAINT input. Add an extra (fourth) column specifying the type of control:
empty, bira-flag, bira-gfp, etc. This file should be trimmed to remove any non-control samples.

```
128_7909	7909_BirAFLAG_April2017	C	bira-flag
128_9522	9522_BirAFLAG_April2017	C	bira-flag
128_8301	8301_FLAG_alone	C	empty
128_8310	8310_FLAG_alone	C	empty
```

### inter

`inter.dat` file file from SAINT input.

```
128_7909	7909_BirAFLAG_April2017	BirA_R118G_H0QFJ5	1075	25
128_7909	7909_BirAFLAG_April2017	NP_000029.2	5	4
128_7909	7909_BirAFLAG_April2017	NP_000048.1	4	4
128_7909	7909_BirAFLAG_April2017	NP_000108.1	7	6
```

### prey

`prey.dat` file file from SAINT input.

```
BirA_R118G_H0QFJ5	321	BirA_R118G_H0QFJ5
NP_000009.1	655	ACADVL
NP_000010.1	427	ACAT1
NP_000091.1	98	CSTB
NP_000099.2	509	DLD
NP_000131.2	423	FECH
NP_000150.1	438	GCDH
NP_000173.2	763	HADHA
NP_000217.2	623	KRT9
NP_000242.1	934	MSH2
```

## Output
* `control-preys`: list of preys found in controls and a breakdown of their abundance
between the different types of controls