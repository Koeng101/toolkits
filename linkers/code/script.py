import numpy as np
import pandas as pd

coreprefix = "AGGTCTCNprefixNNNNNNNNNNNNATTCCGGTCTCTnnnnsuffixCGAGACCA"
coresuffix = "AGGTCTCNprefixnnnnAGAGACCGGAATNNNNNNNNNNNNsuffixCGAGACCA"
def reverse_complement(seq):
    '''Creates reverse complement of a DNA string.'''
    seq = seq.upper()
    return seq.translate(str.maketrans("ATGCN","TACGN"))[::-1]


biobricks_enzymes = {
        "EcorI":"GAATTC",
        "XbaI":"TCTAGA",
        "SpeI":"ACTAGT",
        "PstI":"CTGCAG",
        "NotI":"GCGGCCGC"}
biobricks = [v for k,v in biobricks_enzymes.items()]
homopolymers = ["AAAAA","TTTTT","GGGGG","CCCCC"]
banned = ["GGTCTC","GAAGAC","CGTCTC","ACCTGC","GCGATG","GCTCTTC"] + biobricks + homopolymers
banned = banned + [reverse_complement(x) for x in banned]

def random_dna_sequence(length):
    return ''.join(np.random.choice(('A', 'C', 'T', 'G')) for _ in range(length))

def barcode(length:int=9,banned:list=banned, silent:bool=True)-> str:
    '''Creates a DNA barcode consisting of an enzyme flanked by two equal length random dna sequences that do not contain any banned sequences'''
    full_barcode = random_dna_sequence(length)
    for ban in banned:
        if ban in full_barcode:
            if silent != True:
                print('FOUND {}'.format(ban))
            return barcode(length=length)
    return full_barcode

def chunks(l, n):
    n = max(1, n)
    return (l[i:i+n] for i in range(0, len(l), n))

overhangs = "AAAA GGAG TACT AATG GCTT CGCT ATAG ATGA AGAC AGGT ACAA ACTC TAAC TAGA TTCA TGTG TCGG TCCC GACC GCCG CAAG GTGG TTTC TACG".split(" ")
overhangs = overhangs + [reverse_complement(o) for o in overhangs]

def check(seq):
    b = ["CGTCTC","ACCTGC","GCTCTTC"]
    b = b + [reverse_complement(x) for x in b]
    for x in b:
        if x in seq:
            return False
    return True

def replaceN(seq):
    new = ""
    for char in c:
        if char == "N":
            checkseq = new_core + barcode(1)
            c = 0
            while check(checkseq) == False:
                checkseq = new_core + barcode(1)
                c+=1
                if c > 10:
                    print("OVER 10 at {}".format(checkseq))
            new = checkseq
        else:
            new += char
    return new

def generate_fragments():
    ab = ("AAAA","GGAG")
    ac = ("AAAA","TACT")
    fg = ("CGCT", "ATAG")
    eg = ("GCTT", "ATAG")

    cores = []
    for core,overhang_set in zip([coreprefix,coresuffix], [[ab,ac],[fg,eg]]):
        for prefix,suffix in overhang_set:
            for overhang in overhangs:
                if overhang != "TCAT" and overhang != reverse_complement("TCAT"):
                    checkseq = core.replace("NNNNNNNNN", barcode(length=9)).replace("nnnn", overhang).replace("prefix",prefix).replace("suffix",suffix)
                    while check(checkseq) == False:
                        checkseq = core.replace("NNNNNNNNN", barcode(length=9)).replace("nnnn", overhang).replace("prefix",prefix).replace("suffix",suffix)
                    cores.append("GGAG" + checkseq + "CGCT")
                else:
                    checkseq =core.replace("NNNNNNNNN", barcode(length=9)).replace("nnnn", overhang + "A").replace("prefix",prefix).replace("suffix",suffix)
                    while check(checkseq) == False:
                        checkseq = core.replace("NNNNNNNNN", barcode(length=9)).replace("nnnn", overhang + "A").replace("prefix",prefix).replace("suffix",suffix)
                    cores.append("GGAG" + checkseq + "CGCT")
    
    final_cores = []
    for chunk in chunks(cores,4):
        current_cores = []
        current_cores.append("GTAAAACGACGGCCAGT" + "GCGAAGAC" + "NN" + chunk[0] + "NN" +"GTCTTCNN")
        current_cores.append("CGGCGATG" + "NNNNNNNNNA" + chunk[1] + "NNNNNNNNN" + "ACATCGCNN")
        current_cores.append("ATGAAGAC" + "NN" + chunk[2] + "NN" +"GTCTTCNN")
        current_cores.append("TAGCGATG" + "NNNNNNNNNA" + chunk[3] + "NNNNNNNNN" + "ACATCGC" + "GTCATAGCTGTTTCCTG")
        x = []
        for c in current_cores:
            new_core = ""
            for char in c:
                if char == "N":
                    checkseq = new_core + barcode(1)
                    c = 0
                    while check(checkseq) == False:
                        checkseq = new_core + barcode(1)
                        c+=1
                        if c > 10:
                            print("OVER 10 at {}".format(checkseq))
                    new_core = checkseq
                else:
                    new_core += char
            x.append(new_core)
        final_cores.append(''.join(x))
    for c in final_cores:
        if c.count("GGTCTC") + c.count("GAGACC") != 12:
            print("Found GGTCTC")
            print(c)
            return generate_fragments()
        b = ["CGTCTC","ACCTGC","GCTCTTC"]
        b = b + [reverse_complement(x) for x in b]
        for seq in ["GCGATG", "GAAGAC", reverse_complement("GCGATG"), reverse_complement("GCGATG")]:
            if c.count(seq) != 2:
                print("Found {} {}".format(seq, c.count(seq)))
                print(c)
                return generate_fragments()
        for seq in b:
            if c.count(seq) != 0:
                print(c)
                return generate_fragments()
    return final_cores


overhang_names = []
overhang_seqs = []
for i,overhang in enumerate(overhangs[0:24]):
    overhang_names.append([x for x in range(1,25)][i])
    overhang_seqs.append(overhang)
    overhang_names.append([-x for x in range(1,25)][i])
    overhang_seqs.append(reverse_complement(overhang))
    
print("Building fragments for the following overhangs:")
print(overhangs)
print("Output will be fragments.csv and overhangs.csv")
final_cores = generate_fragments()
pd.DataFrame.from_dict({"fragment_name": ["core_{}-{}".format(i*4, (i*4)+3) for i,core in enumerate(final_cores)], "sequence": final_cores}).to_csv("fragments.csv")
pd.DataFrame.from_dict({"name": overhang_names, "sequence": overhang_seqs}).to_csv("overhangs.csv")

for c in final_cores:
    if c.count("GGTCTC") + c.count("GAGACC") != 12:
        print("FAIL GGTCTC")
