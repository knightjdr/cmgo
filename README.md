# Cellmap analysis modules

Cellmap analysis modules written in GO. This was built with GO modules.

## Build

```
git clone https://github.com/knightjdr/cmgo.git
cd cmgo
go build
```

## Run an analysis module

An options file in JSON format can be used to define run time arguments. Arguments can also be supplied on the command line (these will override corresponding arguments specified in the options file).

A module name must always be supplied for analysis.

```
cmgo -options="config.json"
```

Module specific parameters can be found in the README files in the corresponding folder.

## Enrichment

### Module name: `enrichment-heatmap`

* folder `enrichment/heatmap`

Creates a heat map visualizing region (domain/motif) fold enrichment (scaled to Log2) across NMF or SAFE compartments.

## NMF

### Module name: `nmf-subset`

* folder `nmf/enrichment`

> Subset an NMF basis (prey) matrix to only include preys enriched on in specified compartments.

## Organelle comparison

### Module name: `organelle-overlap`

* folder `organelle/overlap`

Takes two lists of proteins and a txt file with similarity scores between proteins and outputs metrics on the similarity within and between the lists.

### Module name: `organelle-sharedregion`

* folder `organelle/shared`

Takes two lists of proteins and outputs metrics on the regions (domains/motifs) shared between common preys.
