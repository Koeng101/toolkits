
## Give me the pathway
Ok. Here it is, with EC numbers:

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


## EM - Energy production

`formate -> NADH (energy)`

In the EM module, formate is directly converted to NADH for use as energy in the cell.

This single reaction is performed by formate dehydrogenase, creating NADH, which can be used with NADH dehydrogenase to create a proton gradient and generate energy. 

```
EC 1.17.1.9
formate + NAD+ = CO2 + NADH
```

# C1 - Formate conversion
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

# C2 - Glycine cleavage system / synthase

`methylene-THF + CO2 + NH3 -> glycine`

In the C2 module, 5,10-methylene-THF is converted into glycine. Glycine can be used as an amino acid, or given to the C3 module for further conversion.

There is a fantastic wikipedia page on the glycine cleavage system[1], so below will be a focused simplification. The most important part to remember about the glycine cleavage system is that its reactions are reversible, and so direction of metabolic flux is determined by the concentrations of substrates and products within the reaction. For the usage of biosynthesis from formate, we will be running the glycine cleavage system in the direction of glycine synthesis by adding more substrates. 

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

# C3 - Serine and pyruvate synthesis

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

[1] https://en.wikipedia.org/wiki/Glycine_cleavage_system
