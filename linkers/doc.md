# The Linkers Collection

There have been many standard methods for assemblying DNA, with the most notable being the [BioBrick assembly](https://en.wikipedia.org/wiki/BioBrick) method being developed in 2003. However, BioBrick assembly could only assemble 2 DNA parts at once, and so limited the amount of assembly that one could do in a short amount of time. In 2008, a new method called [GoldenGate assembly](https://dx.doi.org/10.1371%2Fjournal.pone.0003647) was developed that overcame limitations of BioBrick assembly, allowing many fragments to be put together at once in a single tube. In 2011, this GoldenGate assembly was standardized with the [MoClo assembly](https://doi.org/10.1371/journal.pone.0016765) method. 

The Linkers Collection is a collection of linkers ([here](https://doi.org/10.3389/fbioe.2019.00271) is a good introduction) for MoClo assembly. These linkers have 2 special attributes:

- The BsaI overhangs are optimized using [empirical data](https://doi.org/10.1371/journal.pone.0238592)
- The assembly is recursive - the same linkers are used at each level of assembly. This is accomplished using methyltransferases which [methylate at CCGG positions](http://www.greatlakesbiotech.org/news/2016/8/26/designing-a-low-cost-molecular-biology-platform)

## How were the linkers designed?

TL;DR: I took the empirically most efficient overhangs and based linkers off of them

In the `./code/` directory, you will find code I used to build these overhangs. The efficiency designer is based off of [this datasheet](https://doi.org/10.1371/journal.pone.0238592.s001) from "Enabling one-pot Golden Gate assemblies of unprecedented complexity using data-optimized assembly design". I wrote this code in Golang, as I eventually want to merge the overhang designer into the [Poly project](https://github.com/TimothyStiles/poly). It also ran very slowly on python. Once I got satisfactory results in Go, I used python (in script.py) to design the particular linker sequences. 4 were placed into each fragment for synthesis, since the minimal sequence for more affordable fragment synthesis providers is ~300bp. Most of the script was created with trial and error to get the right sequences needed for the linker sequences. It works - and only needs to work once. 

## What are linkers?

When building a construct using GoldenGate, simply ligate linkers between your vector and your genes during a GoldenGate reaction to enable use of that gene in multi-gene constructs. 

Typically, you will do an assembly reaction (also known as a level 1 cloning reaction in MoClo lingo) to give context to your gene. For example, you may have a protein called GFP that you wish to express. In this case, you would do a level 1 cloning reaction to contextualize GFP with a proper promoter and terminator for your target organism to make a transcriptional unit (TU). In that reaction, you may have to add linkers to connect your construct into the vector it belongs in. 

Afterwards, you can combine the GFP transcriptional unit to up to 24 other constructs with clever usage of linkers. To answer specifically which ones to use and when, read below.

## What linkers are included in the Linkers Collection?

The "Linkers in Collection" table has all 192 linkers that are included in this collection. There are 96 linkers for building independent transcriptional units with single proteins and 96 linkers for building operons. For both of those categories, linkers are split into 48 prefix and 48 suffix linkers (for the prefix + suffix of any given construct). Those 48 linkers are broken into 24 positive and 24 negative linkers. The positive linkers are used for constructing genes in the forward direction, and negative linkers are used to construct genes in the reverse direction.

Each overhang is assigned a number. The reverse complement of each overhang is represented as the negative version of its number, which is also how we are able to reverse constructs.  

Linkers are named with a 3 number + 1 letter scheme, separated by underscores ( _ ) in the format X_X_XY. The first 2 numbers describe the two overhangs which the linker itself will be cut out with, the third number describes the overhang which the linker will introduce to the construct, and the final letter (F or R) describes the direction, forward or reverse. After a GoldenGate assembly and transformation, anything between the forward and reverse linkers can be cut out with BsaI and used in another assembly. 

For example:
```
1( ----> )2 + -3( --> )-2 + 3( -> )4 = ( ----> <-- -> )
```

## How can you use BsaI in two reactions?

The 3rd internal BsaI site is actually CCGGTCTC instead of only GGTCTC. While BsaI will cut at both sequences, you can methylate CCGG with [HpaII or MspI](http://www.greatlakesbiotech.org/news/2016/8/26/designing-a-low-cost-molecular-biology-platform) and prevent the cutting at CCGGTCTC sites. After assembly and transformation, the CCGGTCTC site will no longer be replicated in the presence of HpaII or MspI, therefore allowing it to be cut.

To initially get methylation at the CCGG sites, you either need to express HpaII methyltransferase within the cell line containing the linker plasmids, or methylate using purified enzymes (for example, from [NEB](https://www.neb.com/products/m0214-hpaii-methyltransferase). 

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
| 1_2_1F |
| 1_2_-1F |
| 1_2_2F |
| 1_2_-2F |
| 1_2_3F |
| 1_2_-3F |
| 1_2_4F |
| 1_2_-4F |
| 1_2_5F |
| 1_2_-5F |
| 1_2_6F |
| 1_2_-6F |
| 1_2_7F |
| 1_2_-7F |
| 1_2_8F |
| 1_2_-8F |
| 1_2_9F |
| 1_2_-9F |
| 1_2_10F |
| 1_2_-10F |
| 1_2_11F |
| 1_2_-11F |
| 1_2_12F |
| 1_2_-12F |
| 1_2_13F |
| 1_2_-13F |
| 1_2_14F |
| 1_2_-14F |
| 1_2_15F |
| 1_2_-15F |
| 1_2_16F |
| 1_2_-16F |
| 1_2_17F |
| 1_2_-17F |
| 1_2_18F |
| 1_2_-18F |
| 1_2_19F |
| 1_2_-19F |
| 1_2_20F |
| 1_2_-20F |
| 1_2_21F |
| 1_2_-21F |
| 1_2_22F |
| 1_2_-22F |
| 1_2_23F |
| 1_2_-23F |
| 1_2_24F |
| 1_2_-24F |
| 1_3_1F |
| 1_3_-1F |
| 1_3_2F |
| 1_3_-2F |
| 1_3_3F |
| 1_3_-3F |
| 1_3_4F |
| 1_3_-4F |
| 1_3_5F |
| 1_3_-5F |
| 1_3_6F |
| 1_3_-6F |
| 1_3_7F |
| 1_3_-7F |
| 1_3_8F |
| 1_3_-8F |
| 1_3_9F |
| 1_3_-9F |
| 1_3_10F |
| 1_3_-10F |
| 1_3_11F |
| 1_3_-11F |
| 1_3_12F |
| 1_3_-12F |
| 1_3_13F |
| 1_3_-13F |
| 1_3_14F |
| 1_3_-14F |
| 1_3_15F |
| 1_3_-15F |
| 1_3_16F |
| 1_3_-16F |
| 1_3_17F |
| 1_3_-17F |
| 1_3_18F |
| 1_3_-18F |
| 1_3_19F |
| 1_3_-19F |
| 1_3_20F |
| 1_3_-20F |
| 1_3_21F |
| 1_3_-21F |
| 1_3_22F |
| 1_3_-22F |
| 1_3_23F |
| 1_3_-23F |
| 1_3_24F |
| 1_3_-24F |
| 6_7_1R |
| 6_7_-1R |
| 6_7_2R |
| 6_7_-2R |
| 6_7_3R |
| 6_7_-3R |
| 6_7_4R |
| 6_7_-4R |
| 6_7_5R |
| 6_7_-5R |
| 6_7_6R |
| 6_7_-6R |
| 6_7_7R |
| 6_7_-7R |
| 6_7_8R |
| 6_7_-8R |
| 6_7_9R |
| 6_7_-9R |
| 6_7_10R |
| 6_7_-10R |
| 6_7_11R |
| 6_7_-11R |
| 6_7_12R |
| 6_7_-12R |
| 6_7_13R |
| 6_7_-13R |
| 6_7_14R |
| 6_7_-14R |
| 6_7_15R |
| 6_7_-15R |
| 6_7_16R |
| 6_7_-16R |
| 6_7_17R |
| 6_7_-17R |
| 6_7_18R |
| 6_7_-18R |
| 6_7_19R |
| 6_7_-19R |
| 6_7_20R |
| 6_7_-20R |
| 6_7_21R |
| 6_7_-21R |
| 6_7_22R |
| 6_7_-22R |
| 6_7_23R |
| 6_7_-23R |
| 6_7_24R |
| 6_7_-24R |
| 5_7_1R |
| 5_7_-1R |
| 5_7_2R |
| 5_7_-2R |
| 5_7_3R |
| 5_7_-3R |
| 5_7_4R |
| 5_7_-4R |
| 5_7_5R |
| 5_7_-5R |
| 5_7_6R |
| 5_7_-6R |
| 5_7_7R |
| 5_7_-7R |
| 5_7_8R |
| 5_7_-8R |
| 5_7_9R |
| 5_7_-9R |
| 5_7_10R |
| 5_7_-10R |
| 5_7_11R |
| 5_7_-11R |
| 5_7_12R |
| 5_7_-12R |
| 5_7_13R |
| 5_7_-13R |
| 5_7_14R |
| 5_7_-14R |
| 5_7_15R |
| 5_7_-15R |
| 5_7_16R |
| 5_7_-16R |
| 5_7_17R |
| 5_7_-17R |
| 5_7_18R |
| 5_7_-18R |
| 5_7_19R |
| 5_7_-19R |
| 5_7_20R |
| 5_7_-20R |
| 5_7_21R |
| 5_7_-21R |
| 5_7_22R |
| 5_7_-22R |
| 5_7_23R |
| 5_7_-23R |
| 5_7_24R |
| 5_7_-24R |

