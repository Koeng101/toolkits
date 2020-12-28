package main

import (
	"fmt"
	"github.com/koeng101/poly"
	"strconv"
	"strings"
)

var BannedSeq = []string{"GGTCTC", "GAGACC", "CGTCTC", "GAGACG", "GAAGAC", "GTCTTC", "GCGATG", "CATCGC", "GCTCTTC", "GAAGAGC", "ACCTGC", "GCAGGT"}

func buildSeq(fasta poly.Fasta) poly.Fasta {
	nameStrList := strings.Split(fasta.Name, "|")
	overhangs := strings.Split(nameStrList[1], ",")

	name := nameStrList[0]
	overhangFor := overhangs[0]
	overhangRev := overhangs[1]

	newSeq := "GGAGGTCTCA" + overhangFor + strings.ToUpper(fasta.Sequence) + overhangRev + "AGAGACCGCT"

	return poly.Fasta{name, newSeq}

}

func main() {
	// Parse those FASTA files
	fastasChan := make(chan poly.Fasta)
	go poly.ReadFASTAConcurrent("genes.fasta", fastasChan)
	var fastas []poly.Fasta
	for fasta := range fastasChan {
		fastas = append(fastas, fasta)
	}

	// Build new fasta files with the proper overhangs and such
	var newFastas []poly.Fasta
	for _, fasta := range fastas {
		newFastas = append(newFastas, buildSeq(fasta))
	}

	//fmt.Println(len(newFastas))
	// There are 10 sequences in that file. Most the vectors I plan on building right now
	// are for E.coli, and so they don't need the extra spots that are reserved for
	// homology or target organism selections / origins. Therefore, I'll add in a few barcodes,
	// attached to the ends of the genes that I actually want to synthesize.
	var barcodes []string
	barcodesChan := make(chan string)
	go poly.UniqueSequence(barcodesChan, 24, 4, BannedSeq, []func(string) bool{})
	for barcode := range barcodesChan {
		barcodes = append(barcodes, barcode)
	}

	var finalFastas []poly.Fasta
	// Append the barcodes to the newFasta files
	for i := 0; i < 5; i++ {
		finalFastas = append(finalFastas, poly.Fasta{newFastas[i].Name + "__" + "vecForSpacer-" + strconv.Itoa(i), "GAAGACGA" + newFastas[i].Sequence + "GAGTCTTC" + "GCGATGGGTAGCCTCAGGAGGTCTCA" + "ATAG" + barcodes[i] + "ATGA" + "AGAGACCGCTGTACGCCTCACATCGC"})
	}
	for i := 5; i < 10; i++ {
		finalFastas = append(finalFastas, poly.Fasta{newFastas[i].Name + "__" + "vecRevSpacer-" + strconv.Itoa(i-5), "GAAGACGA" + newFastas[i].Sequence + "GAGTCTTC" + "GCGATGGGTAGCCTCAGGAGGTCTCA" + "ACAA" + barcodes[i] + "AAAA" + "AGAGACCGCTGTACGCCTCACATCGC"})
	}

	errMsg := "No errors."
	for _, fasta := range finalFastas {
		switch {
		case strings.Count(fasta.Sequence, "GGTCTC") != 2:
			errMsg = "Found more than 2 GGTCTC"
		case strings.Count(fasta.Sequence, "GAGACC") != 2:
			errMsg = "Found more than 2 GAGACC"
		case strings.Count(fasta.Sequence, "GAAGAC") != 1:
			errMsg = "Found more than 1 GAAGAC"
		case strings.Count(fasta.Sequence, "GTCTTC") != 1:
			errMsg = "Found more than 1 GTCTTC"
		case strings.Count(fasta.Sequence, "GCGATG") != 1:
			errMsg = "Found more than 1 GCGATG"
		case strings.Count(fasta.Sequence, "CATCGC") != 1:
			errMsg = "Found more than 1 CATCGC"
		}
	}
	// Read out those names
	if errMsg == "No errors." {
		fmt.Println("Name, Sequence")
		for _, fasta := range finalFastas {
			fmt.Println(fasta.Name + "," + fasta.Sequence)
		}
	} else {
		fmt.Println(errMsg)
	}
}
