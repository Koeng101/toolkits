# The Basics

The basics toolkit has the raw basic genes needed to make nice multi-gene cassettes. It contains only a few genes, listed below:

1. Origins: colE1, p15a, cloDF13, pSB2K3-ori
2. Positive selection: AmpR, KanR, CamR, SpecR
3. Negative selection: ccdB, sacB
4. Packaging: oriT

The designs for this toolkit primarily come from Keoni Gandall's previous work to produce vectors for OpenEnzyme. Many of those designs were derived from SEVA parts. 

Sizes:
 - 546 p15a
 - 589 colE1
 - 739 cloDF13
 - ~3000 pSB2K3-ori
 - 1039 ampR
 - 927 kanR
 - 783 camR
 - 989 specR
 - 724 ccdB
 - 1482 sacB
 - 246 oriT

Stuff I plan on PCRing (assuming $10 per pair, minimum $100 of primers):
- colE1
- pSB2K3-ori
- ampR
- ccdB
- sacB
- [kanR](https://stanford.freegenes.org/collections/open-genes/products/expression-tookit)

In total, that puts the synthesis cost at about ~$500 for all of these genes, plus the cloning costs. 

## Design

I've included 4 origins. ColE1 and p15a are included because they are the most popular high and medium copy vectors, but I also included CloDF13 and pSB2K3-ori. [CloDF13](https://openwetware.org/wiki/CH391L/S12/Origins_of_Replication) has about ~2-3x the copy number of p15a and is a compatible origin, meaning a CloDF13 and p15a origin can coexist in a cell happily. I've used this origin before, and it has worked for me. I am not including pSC101 in this basic set, since derivatives are often used in applications with heat curing, and since it is [rolling circle replicating](https://en.wikipedia.org/wiki/Rolling_circle_replication) based rather than [theta](https://en.wikipedia.org/wiki/Theta_structure) based. Oftentimes, this means there is a complementing protein that is usually expressed alongside the origin - because of this, pSC101 is more related to the R6K conditional origin than the theta p15a, CloDF13, or ColE1 origins. I plan on synthesizing pSC101 and R6K separately later. 

pSB2K3 inducible copy vector is very useful for large DNA molecules.

I did not include Tet in the basic positive selection markers because it is light sensitive, and I don't like handling tubes wrapped in tin foil. I don't have any data on how much this actually matters, but I don't want to use it right now, so I am not paying for its synthesis. I will probably go back and add it later. 

ccdB and sacB are for building prebuilt vectors - this is useful for when you use the same vector over and over, which honestly, will be pretty often. 

In order to not have conflicting selections, the markers will be split up, and you'll have to reassemble them using GoldenGate at build time. This is so I don't have to manage the complexity of negative selection on builds with those parts. Easier to add 2 different parts than to have sucrose and non-sucrose plates (if I wanted to build a ccdB prebuilt vector from sacB-resistance marker plasmids).


