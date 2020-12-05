# Synthesis Plan


## EM module

In ["Growth of E. coli on formate and methanol via the reductive glycine pathway"](https://doi.org/10.1038/s41589-020-0473-5) the EM module was built from the gene fdh (formate dehydrogenase) in Pseudomonas sp. (strain 101). We will be using that exact gene, codon optimized for Bacillus subtilis. 

- [fdh](https://www.uniprot.org/uniprot/P33160)

## C1 module

In ["Growth of E. coli on formate and methanol via the reductive glycine pathway"](https://doi.org/10.1038/s41589-020-0473-5) the C1 module was built from genes from Methylobacterium extorquens - in particular, fhs (formate--tetrahydrofolate ligase), fchA (methenyltetrahydrofolate cyclohydrolase), and mtdA (methylenetetrahydrofolate dehydrogenase). We'll use these exact genes, codon optimized for Bacillus subtilis.

- [fhs](https://www.uniprot.org/uniprot/Q83WS0)
- [fchA](https://www.uniprot.org/uniprot/Q49135)
- [mtdA](https://www.uniprot.org/uniprot/P55818)

## C2 module

The C2 module is made of the genes involved in the glycine cleavage system, which exists in both Escherichia coli and Bacillus subtilis. Therefore, the 3-4 genes necessary for this pathway will simply be PCR'd out of Escherichia coli and Bacillus subtilis.

## C3 module

The C3 module is also made out of genes naturally occuring in Escherichia coli and Bacillus subtilis. These too will be PCR'd out.

## Deletions

The deletions are a bit more tricky. There are many great systems for doing genomic modification of Bacillus subtilis, but we will be using CRISPR plasmids from "[New CRISPR-Cas9 vectors for genetic modifications of Bacillus species](https://doi.org/10.1093/femsle/fny284)". Since we only need to do 4 replacements, we will just use oligos with the existing protocol in that paper. If it works, we will move into much higher throughput work to target all genes for deletions. The primers necessary to do this are part of primers.fasta.
