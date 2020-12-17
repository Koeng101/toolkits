# The Linkers Plate

There have been many standard methods for assemblying DNA, with the most notable being the [BioBrick assembly](https://en.wikipedia.org/wiki/BioBrick) method being developed in 2003. However, BioBrick assembly could only assemble 2 DNA parts at once, and so limited the amount of assembly that one could do in a short amount of time. In 2008, a new method called [GoldenGate assembly](https://dx.doi.org/10.1371%2Fjournal.pone.0003647) was developed that overcame limitations of BioBrick assembly, allowing many fragments to be put together at once in a single tube. In 2011, this GoldenGate assembly was standardized with the [MoClo assembly](https://doi.org/10.1371/journal.pone.0016765) method. 

The Connector's Collection is a collection of linkers/connectors ([here](https://doi.org/10.3389/fbioe.2019.00271) is a good introduction to the idea of linkers/connectors) for MoClo assembly. These linkers have 2 special attributes:

- The BsaI overhangs are optimized using [empirical data](https://doi.org/10.1371/journal.pone.0238592)
- The assembly is recursive - the same linkers are used at each level of assembly. This is accomplished using methyltransferases which [methylate at CCGG positions](http://www.greatlakesbiotech.org/news/2016/8/26/designing-a-low-cost-molecular-biology-platform)

## How were the linkers designed?

TL;DR: I took the empirically most efficient overhangs and based linkers off of them

In the `./code/` directory, you will find code I used to build these overhangs. The efficiency designer is based off of [this datasheet](https://doi.org/10.1371/journal.pone.0238592.s001) from "Enabling one-pot Golden Gate assemblies of unprecedented complexity using data-optimized assembly design". I wrote this code in Golang, as I eventually want to merge the overhang designer into the [Poly project](https://github.com/TimothyStiles/poly). It also ran very slowly on python. Once I got satisfactory results in Go, I used python (in script.py) to design the particular linker sequences. 4 were sequenced into each fragment for synthesis, since the minimal sequence for more affordable fragment synthesis providers is ~300bp. Most of the script was created with trial and error to get the right sequences needed for the linker sequences. It works - and only needs to work once. 

## What are linkers?

When building a construct using GoldenGate, simply ligate linkers between your vector and your genes during a GoldenGate reaction to enable use of that gene in multi-gene constructs. 

Typically, you will do an assembly reaction (also known as a level 1 cloning reaction in MoClo lingo) to give context to your gene. For example, you may have a protein called GFP that you wish to express. In this case, you would do a level 1 cloning reaction to contextualize GFP with a proper promoter and terminator for your target organism to make a transcriptional unit (TU). In that reaction, you may have to add linkers to connect your construct into the vector it belongs in. 

Afterwards, you can combine the GFP transcriptional unit to up to 24 other constructs with clever usage of linkers. To answer specifically which ones to use and when, read below.

## What linkers are included in the Linkers Collection?

The "Linkers in Collection" table has all 192 linkers that are included in this collection. There are 96 linkers for building independent transcriptional units with single proteins and 96 linkers for building operons. For both of those categories, linkers are split into 48 prefix and 48 suffix linkers (for the prefix + suffix of any given construct). Those 48 linkers are broken into 24 positive and 24 negative linkers. The positive linkers are used for constructing genes in the forward direction, and negative linkers are used to construct genes in the reverse direction.

Each overhang is assigned a number. The reverse complement of each overhang is represented as the negative version of its number, which is also how we are able to reverse constructs.  

For example:
```
1( ----> )2 + -3( --> )-2 + 3( -> )4 = ( ----> <-- -> )
```

# Overhang table

| Name | Sequence | 
| - | - |
| 1 | AAAA |
| -1 | TTTT |
| 2 | GGAG |
| -2 | CTCC |
| 3 | TACT |
| -3 | AGTA |
| 4 | AATG |
| -4 | CATT |
| 5 | GCTT |
| -5 | AAGC |
| 6 | CGCT |
| -6 | AGCG |
| 7 | ATAG |
| -7 | CTAT |
| 8 | ATGA |
| -8 | TCAT |
| 9 | AGAC |
| -9 | GTCT |
| 10 | AGGT |
| -10 | ACCT |
| 11 | ACAA |
| -11 | TTGT |
| 12 | ACTC |
| -12 | GAGT |
| 13 | TAAC |
| -13 | GTTA |
| 14 | TAGA |
| -14 | TCTA |
| 15 | TTCA |
| -15 | TGAA |
| 16 | TGTG |
| -16 | CACA |
| 17 | TCGG |
| -17 | CCGA |
| 18 | TCCC |
| -18 | GGGA |
| 19 | GACC |
| -19 | GGTC |
| 20 | GCCG |
| -20 | CGGC |
| 21 | CAAG |
| -21 | CTTG |
| 22 | GTGG |
| -22 | CCAC |
| 23 | TTTC |
| -23 | GAAA |
| 24 | TACG |
| -24 | CGTA |

# Linkers in Collection

| Name |
| - |
| 1_2_1 |
| 1_2_-1 |
| 1_2_2 |
| 1_2_-2 |
| 1_2_3 |
| 1_2_-3 |
| 1_2_4 |
| 1_2_-4 |
| 1_2_5 |
| 1_2_-5 |
| 1_2_6 |
| 1_2_-6 |
| 1_2_7 |
| 1_2_-7 |
| 1_2_8 |
| 1_2_-8 |
| 1_2_9 |
| 1_2_-9 |
| 1_2_10 |
| 1_2_-10 |
| 1_2_11 |
| 1_2_-11 |
| 1_2_12 |
| 1_2_-12 |
| 1_2_13 |
| 1_2_-13 |
| 1_2_14 |
| 1_2_-14 |
| 1_2_15 |
| 1_2_-15 |
| 1_2_16 |
| 1_2_-16 |
| 1_2_17 |
| 1_2_-17 |
| 1_2_18 |
| 1_2_-18 |
| 1_2_19 |
| 1_2_-19 |
| 1_2_20 |
| 1_2_-20 |
| 1_2_21 |
| 1_2_-21 |
| 1_2_22 |
| 1_2_-22 |
| 1_2_23 |
| 1_2_-23 |
| 1_2_24 |
| 1_2_-24 |
| 1_3_1 |
| 1_3_-1 |
| 1_3_2 |
| 1_3_-2 |
| 1_3_3 |
| 1_3_-3 |
| 1_3_4 |
| 1_3_-4 |
| 1_3_5 |
| 1_3_-5 |
| 1_3_6 |
| 1_3_-6 |
| 1_3_7 |
| 1_3_-7 |
| 1_3_8 |
| 1_3_-8 |
| 1_3_9 |
| 1_3_-9 |
| 1_3_10 |
| 1_3_-10 |
| 1_3_11 |
| 1_3_-11 |
| 1_3_12 |
| 1_3_-12 |
| 1_3_13 |
| 1_3_-13 |
| 1_3_14 |
| 1_3_-14 |
| 1_3_15 |
| 1_3_-15 |
| 1_3_16 |
| 1_3_-16 |
| 1_3_17 |
| 1_3_-17 |
| 1_3_18 |
| 1_3_-18 |
| 1_3_19 |
| 1_3_-19 |
| 1_3_20 |
| 1_3_-20 |
| 1_3_21 |
| 1_3_-21 |
| 1_3_22 |
| 1_3_-22 |
| 1_3_23 |
| 1_3_-23 |
| 1_3_24 |
| 1_3_-24 |
| 6_7_1 |
| 6_7_-1 |
| 6_7_2 |
| 6_7_-2 |
| 6_7_3 |
| 6_7_-3 |
| 6_7_4 |
| 6_7_-4 |
| 6_7_5 |
| 6_7_-5 |
| 6_7_6 |
| 6_7_-6 |
| 6_7_7 |
| 6_7_-7 |
| 6_7_8 |
| 6_7_-8 |
| 6_7_9 |
| 6_7_-9 |
| 6_7_10 |
| 6_7_-10 |
| 6_7_11 |
| 6_7_-11 |
| 6_7_12 |
| 6_7_-12 |
| 6_7_13 |
| 6_7_-13 |
| 6_7_14 |
| 6_7_-14 |
| 6_7_15 |
| 6_7_-15 |
| 6_7_16 |
| 6_7_-16 |
| 6_7_17 |
| 6_7_-17 |
| 6_7_18 |
| 6_7_-18 |
| 6_7_19 |
| 6_7_-19 |
| 6_7_20 |
| 6_7_-20 |
| 6_7_21 |
| 6_7_-21 |
| 6_7_22 |
| 6_7_-22 |
| 6_7_23 |
| 6_7_-23 |
| 6_7_24 |
| 6_7_-24 |
| 5_7_1 |
| 5_7_-1 |
| 5_7_2 |
| 5_7_-2 |
| 5_7_3 |
| 5_7_-3 |
| 5_7_4 |
| 5_7_-4 |
| 5_7_5 |
| 5_7_-5 |
| 5_7_6 |
| 5_7_-6 |
| 5_7_7 |
| 5_7_-7 |
| 5_7_8 |
| 5_7_-8 |
| 5_7_9 |
| 5_7_-9 |
| 5_7_10 |
| 5_7_-10 |
| 5_7_11 |
| 5_7_-11 |
| 5_7_12 |
| 5_7_-12 |
| 5_7_13 |
| 5_7_-13 |
| 5_7_14 |
| 5_7_-14 |
| 5_7_15 |
| 5_7_-15 |
| 5_7_16 |
| 5_7_-16 |
| 5_7_17 |
| 5_7_-17 |
| 5_7_18 |
| 5_7_-18 |
| 5_7_19 |
| 5_7_-19 |
| 5_7_20 |
| 5_7_-20 |
| 5_7_21 |
| 5_7_-21 |
| 5_7_22 |
| 5_7_-22 |
| 5_7_23 |
| 5_7_-23 |
| 5_7_24 |
| 5_7_-24 |

