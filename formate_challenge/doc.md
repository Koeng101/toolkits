
# The Bacillus Formate Bioeconomy docs
This toolkit contains the specific genes necessary to give Bacillus subtilis synthetic formatotrophy. A formatotrophic Bacillus subtilis isn't supplied - that is up to you! If you need a good review on this system in Escherichia coli, check out [Growth of E. coli on formate and methanol via the reductive glycine pathway](https://doi.org/10.1038/s41589-020-0473-5). The pathway described there can be a little hard to understand, so I've broken down the biosynthesis pathway with EC numbers and explanations below. 

### Why synthetic formatotrophy?
The idea of synthetic formatotrophy is to build a pathway that allows an organism to depend on formate instead of glucose for cellular energy and carbon biomass. This formate can be created with renewable energy sources, removing photosynthesis as the bottom rung of the trophic ladder. If humanity can increase electricity production enough, synthetic formatotrophy opens up a whole new way to create biomass for useful chemicals, food, and more!

There are 2 pathways that use formate to generate biomass in non-formate utilizing organisms - the [rubisco pathway](https://doi.org/10.1016/j.cell.2019.11.009), or the [reductive glycine pathway](https://doi.org/10.1038/s41589-020-0473-5). The rubsico pathway uses formate to power a non-native calvin cycle, while the reductive glycine pathway uses formate to synthesize glycine,
serine, and pyruvate. The rubisco pathway's cycle creates a lot of intermediate metabolic products that are used in other pathways, while the reductive glycine pathway is more linear and easi
er to model. This becomes apparent with the length of time necessary to achieve synthetic formatotrophy - with the Rubisco pathway, it took about ~200 days of continuous evolution to get a cell that could live purely on formate, while with the reductive glycine pathway, a formatotroph could be found immediately after all modifications were made. In addition to this advantage, we will be using the reductive glycine pathway since there are auxotrophic methods to test if the formatotrophy pathways are functioning.

### How do you create synthetic formatotrophy?
To get from formate to cell energy, you convert formate to NADH, which can be used for the electron transport chain. To get from formate to biomass, you convert formate to glycine, then to serine, and finally to pyruvate. Pyruvate can be used in central metabolism, while glycine and serine can be used as amino acids in proteins. 

The entire pathway is broken up into 4 "modules" - EM, C1, C2, and C3. Successful expression of EM will allow the cell to use formate for energy. Successful expression of C1 and C2 allow glycine auxotrophs to synthesize glycine. Successful expression of C3 will allow serine auxotrophs to synthesize serine (and pyruvate). In Escherichia coli, there are a total there are 9 genes that need to be expressed: 1 for EM, 3 for C1, 3 for C2, and 2 for C3. The 4 genes in EM and C1 need to be introduced synthetically from different organisms (Pseudomonas strain 101 is used for EM and Methylobacterium extorquens is used for C1) while the 5 genes in C2 and C3 already exist in the current biosynthetic pathways of Escherichia coli, but simply need to be overexpressed. In Bacillus subtilis, the same 4 genes for EM and C1 will need to be introduced synthetically and a different set of genes are used for C2 and C3. 

In Escherichia coli, 6 deletions are required to test the functionality of the formate uptake and conversion pathway - serA, kbl, ltaE, aceA, glyA, and sdaA. serA forces serine auxotrophy, kbl and ltaE remove theorine clevage to glycine, and aceA forces glycine auxotrophy. glyA and sdaA are overexpressed in C3, and so were deleted from the genome (presumably so that homologous recombination wouldn't occur between the native and introduced sequences). In Bacillus subtilis, at least 4 of those deletions would be necessary to test if C1, C2, and C3 are functional. 

Let's do a deep dive of the actual pathway used for synthetic formatotrophy. 

## The Synthetic Formatotrophy Pathway
This is the compressed pathway using EC numbers for synthetic formatotrophy.
```
# Energy 
1.17.1.9

# Biomass
6.3.4.3 -> 3.5.4.9 -> 1.5.1.5 -> 2.1.2.10 -> 1.4.4.2[1][2] -> 2.1.2.1[3] -> 4.3.1.17[4]
[1] Recycles with 1.8.1.4 to enter at 2.1.2.10
[2] Glycine output can be used as amino acid
[3] Serine output can be used as amino acid
[4] Pyruvate output can be used for energy or for fatty acid synthesis
```

### EM - Energy production

`formate -> NADH (energy)`

In the EM module, formate is directly converted to NADH for use as energy in the cell.

This single reaction is performed by formate dehydrogenase, creating NADH, which can be used with NADH dehydrogenase to create a proton gradient and generate energy (ATP). 

```
EC 1.17.1.9 - formate dehydrogenase:
formate + NAD+ = CO2 + NADH
```

### C1 - Formate conversion
`formate -> 5,10-methylene-THF`

In the C1 module, formate is converted to a more familiar molecule, 5,10-methylene-THF, for entry into the C2 module and C3 module. 

These 3 reactions are quite simple and each rely on a single protein. Formate-THF ligase first ligates THF onto the formate, which can then be converted to 5,10-methyl-THF using methenyltetrahydrofolate cyclohydrolase. Methylenetetrahydrofolate dehydrogenase finally converts 5,10-methyl-THF to 5,10-methylene-THF.

```
EC 6.3.4.3 - formate---tetrahydrofolate ligase:
formate + THF + ATP = 10-formyl-THF + ADP + Pi

EC 3.5.4.9 - methenyltetrahydrofolate cyclohydrolase:
10-formyl-THF = H2O, 5,10-methyl-THF

EC 1.5.1.5 - methylenetetrahydrofolate dehydrogenase:
5,10-methyl-THF + NADPH + H+ = 5,10-methylene-THF + NADP+
```

### C2 - Glycine cleavage system / synthase

`5,10-methylene-THF + CO2 -> glycine`

In the C2 module, 5,10-methylene-THF is converted into glycine. Glycine can be used as an amino acid, or given to the C3 module for further conversion.

There is a fantastic wikipedia page on the [glycine cleavage system](https://en.wikipedia.org/wiki/Glycine_cleavage_system), so below will be a focused simplification. The most important part to remember about the glycine cleavage system is that its reactions are reversible, and so direction of metabolic flux is determined by the concentrations of substrates and products within the reaction. For the usage of biosynthesis from formate, we will be running the glycine cleavage system in the direction of glycine synthesis by adding more substrates. 

There are 4 proteins in the system: The T-protein, the P-protein, the L-protein, and the H-protein. The H-protein facilitates transfer of molecules between the 3 other proteins. The T-protein does reaction EC 2.1.2.10, using the H-protein, 5,10-methylene-THF from the C1 system, and ammonia to produce THF and the second H-protein state. The H-protein is then passed to the P-protein, which uses the H-protein and CO2 to produce glycine and the third H-protein state. Finally, the H-protein is passed to the L-protein with some NAD+ to produce NADH, H+, and the original H-protein, which is recycled into the system. 

```
Notes:
- H(x) represents H protein in different states
- 5,10-methylene-THF is the output of C1

EC 2.1.2.10 - T protein:
H(0) + 5,10-methylene-THF + NH3 = H(1) + THF

EC 1.4.4.2 - P protein:
H(1) + CO2 = glycine + H(2)

EC 1.8.1.4 - L protein:
H(2) + NAD+ = H(0) + NADH + H+
```

### C3 - Serine and pyruvate synthesis

`glycine + methylene-THF -> serine + pyruvate`

In the C3 module, glycine is converted into serine and pyruvate to make biomass.

The C3 module uses 2 enzymes, serine hydroxymethylase and serine deaminase. Serine hydroxymethylase converts glycine and 5,10-methylene-THF into serine, which can directly be used as amino acid. However, instead of being used as an amino acid, the serine can also be used by serine deaminase to create pyruvate, which can be used for energy or for creating fatty acids.

```
Notes:
- EC 4.3.1.17, serine deaminase, actually does 3 reactions to get to its final product
- EC 4.3.1.19, threonine deaminase, can work as a serine deaminase in several organisms
- 5,10-methylene-THF is the output of C1 and glycine is the output of C2

EC 2.1.2.1 - serine hydroxymethylase:
5,10-methylene-THF + glycine + H2O = THF + L-serine

EC 4.3.1.17 - serine deaminase:
L-serine = pyruvate + NH3
```

## The Deletions

As said above, there are 4 genes that must be deleted to test the synthetic formatotrophy pathway. Generally, these genes prevent synthesis of serine and glycine from native pathways. Glycine and serine auxotrophy can then be used to test if formatotrophy is functioning in the target organism. Quite a clever trick!

```
Notes:
- serA is the first step of the serine biosynthesis (3-phosphooxypyruvate is used to synthesize serine)
- In literature, it seems glyoxylate synthesized by aceA can be used for glycine synthesis under stress conditions

serA - EC 1.1.1.95
(2R)-3-phosphoglycerate + NAD+ = 3-phosphooxypyruvate + H+ + NADH

kbl - EC 2.3.1.29
(2S)-2-amino-3-oxobutanoate + CoA = glycine + acetyl-CoA

ltaE - EC 4.1.2.48
L-threonine = acetaldehyde + glycine

aceA - EC 4.1.3.1
D-threo-isocitrate = glyoxylate + succinate
```

Gene deletions can be done in several different ways, though the most promising would likely be the [CRISPR plasmids](https://doi.org/10.1093/femsle/fny284). 
