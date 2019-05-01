# Module name: `summary-crapome`

> Generates a CRAPome matrix from SAINT input files

| parameters | definition | default |
|------------|------------|---------|
| baitFiles | bait.dat: files with bait sample IDs, bait name and control status (C: control, T: non-control)  | |
| interactionFiles | inter.dat: files with bait ID and name, prey accession and spectral count  | |
| outFile | output matrix | crapome-matrix.txt |
| preyFiles | prey.dat: files with prey accession, sequence length and gene name | |

Multiple SAINT tasks can be merged by listing multiple bait.dat, inter.dat and prey.dat files separated by a semicolon. Each of these should be listed in a consistent order:
```
-baitFiles="bait-task1.dat;bait-task2.dat" -interactionFiles="inter-task1.dat;inter-task2.dat" -preyFiles="prey-task1.dat;prey-task2.dat"
```

## Example file formats

### baitFile (bait.dat)

header: projectID_sampleID | bait name | control status
```
128_468	ACTB	T
128_492	ACTB	T
128_590	ATP2A1	T
128_546	ATP2A1	T
128_737	737_BirAFLAG	C
128_825	825_BirAFLAG	C
```

### interactionFile (inter.dat)
header: projectID_sampleID | bait name | prey accession | spectral count | unique peptides |
```
128_468	ACTB	BirA_R118G_H0QFJ5	1410	23
128_468	ACTB	NP_000029.2	4	3
128_468	ACTB	NP_000108.1	2	2
128_468	ACTB	NP_000280.1	3	2
128_468	ACTB	NP_000402.3	11	10
128_468	ACTB	NP_000652.2	7	5
```

### preyFile (prey.dat)
header: accession | prey sequence length | prey name |
```
BirA_R118G_H0QFJ5	321	BirA_R118G_H0QFJ5
NP_000029.2	2843	APC
NP_000108.1	254	EMD
NP_000280.1	780	PFKM
NP_000402.3	726	HLCS
NP_000652.2	192	RPL9
NP_000911.2	1178	PC
```

### Output
* txt matrix file formatted for CRAPome with prey summaries across control samples