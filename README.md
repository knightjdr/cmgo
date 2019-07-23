# Cellmap analysis modules

Cellmap analysis modules written in GO. This was built with GO modules.

## Build

```
git clone https://github.com/knightjdr/cmgo.git
cd cmgo
go build ./...
```

## Run an analysis module

An options file in JSON format can be used to define run time arguments. Arguments can also be supplied on the command line (these will override corresponding arguments specified in the options file).

A module name must always be supplied for analysis.

```
cmgo -options="config.json" -module="module-name"
```

Module specific parameters can be found in the README files in the corresponding folder.

## Analysis

Modules for generating files for upload and query file processing.

### Module name: `analysis-dbgenes`

* folder `analysis/dbgenes`

Generates a txt file of genes in FASTA database

## Assessment

### Module name: `bait-gradient`

* folder `assessment/bait/gradient`

> Takes an ordered bait list and expected localizations for the baits and outputs a gradient showing
> adjacent bait similarity: black if the two adjacent baits share an expected localization, grey
> if only one does and white otherwise.

### Module name: `assessment-hydropathy`

* folder `assessment/hydropathy`

> Measures the average hydropathy of proteins in a SAINT file and a BioPlex dataset

### Module name: `assessment-localization`

* folder `assessment/localization/evaluate`

> Report on which localizations are previously known.

## Enrichment

Modules for calculating term enrichment in compartments and associated images.

### Module name: `enrichment-heatmap`

* folder `enrichment/heatmap`

Creates a heat map visualizing region (domain/motif) fold enrichment (scaled to Log2) across NMF or SAFE compartments.

## LBA

### Module name: `lba-correlation`

* folder `lba/correlation`

> Calculate correlation between LBA profiles and output Cytoscape formatted file for viewing network.

### Module name: `lba-enrichment`

* folder `lba/enrichment`

> Enrichment performs GO enrichments of preys from a SAINT file using LBA.

### Module name: `lba-localize`

* folder `lba/localize`

> Localize preys based on enriched terms for LBA

### Module name: `nmf-v-safe`

* folder `assessment/localization/nmfsafe`

> Check concordance between NMF and SAFE localizations

## Network

Modules for creating networks

### Module name: `network-svg`

> Create an svg network for NMF or SAFE

### Module name: `network-tsnecytoscape`

> Convert a tSNE coordinate system into a .cyjs file for Cytoscape

## NMF

Modules for NMF (and related) analysis.

### Module name: `nmf-robustness`

* folder `nmf/robustness`

> Evaluates how robust each rank's GO terms are based on the genes used
> to define it. It will first define the rank by performing a GO enrichment
> using at most maxGenesPerRank, and then repeat this when using 90%, 80%...
> of the maxGenesPerRank. For each evaluation it will calculate the RBD between
> the list generated for maxGenesPerRank and the fractional list.

### Module name: `nmf-subset`

* folder `nmf/subset`

> Subset an NMF basis (prey) matrix to only include preys enriched on in specified compartments.

### Module name: `nmf-uv`

* folder `nmf/uv`

> For all prey genes not used to define an NMF rank, it calculates what
> proportion are being assigned to a previously known localization. UV:
> unused validation

## Organelle comparison

Modules for comparing cell map compartments.

### Module name: `organelle-overlap`

* folder `organelle/overlap`

Takes two lists of proteins and a txt file with similarity scores between proteins and outputs metrics on the similarity within and between the lists.

### Module name: `organelle-sharedregion`

* folder `organelle/shared`

Takes two lists of proteins and outputs metrics on the regions (domains/motifs) shared between common preys.

## Summary

SAINT file summaries and information

### Module name: `summary-crapome`

* folder `summary/crapome`

Generates a CRAPome matrix from SAINT input files

### Module name: `summary-notsignificant`

* folder `summary/notsignificant`

Generates a list of preys that were not significant with any bait in a SAINT file.

# Files for website

1. CRAPome matrix
  * module: `summary-crapome`
  * filename: `crapome-matrix-v1.txt`
  * destination: client/resources/downloads
2. preys detected in samples that are not significant
  * module: `summary-notsignificant`
  * filename: `not-significant-v1.txt`
  * destination: client/resources/downloads
3. sequence database
  * download FASTA database from ProHits
  * filename: `sequence-database-v1.txt`
  * destination: client/resources/downloads
4. genes in database
  * module: `analysis-dbgenes`
  * filename: `v1_dbgenes.txt`
  * destination: api/app/data/genes
