package main

import (
	"fmt"
	"github.com/koeng101/poly"
	"strconv"
)

type Primer struct {
	Name     string
	Sequence string
}

func main() {
	// Generate Standard Primers
	var Primers []Primer
	standardFor := "ATGGGCATTGTGGAGGTAAGC"
	standardRev := "TCAGTAAGTTCGCTCGTCAATGAG"
	oriFor := "GCAGAGCGAGGTATGTAGGC"
	oriRev := "GAGCTATGAGAAAGCGCCAC"

	// Make deBruijn barcodes for standard primers
	standardBarcodes := make(chan string)
	go poly.UniqueSequence(standardBarcodes, 40, 7, []string{"AAAAA", "TTTTT", "GGGGG", "CCCCC"}, []func(string) bool{})
	barcodes := []string{}
	for barcode := range standardBarcodes {
		barcodes = append(barcodes, barcode)
	}
	for i := 0; i < 96; i++ {
		if i < 48 {
			Primers = append(Primers, Primer{"standardFor_" + strconv.Itoa(i), barcodes[i] + standardFor})
		} else {
			Primers = append(Primers, Primer{"standardRev_" + strconv.Itoa(i-48), barcodes[i] + standardRev})
		}
	}

	// Make specific ori primers
	oriBarcodes := make(chan string)
	go poly.UniqueSequence(oriBarcodes, 20, 5, []string{"AAAAA", "TTTTT", "GGGGG", "CCCCC"}, []func(string) bool{})
	barcodes = []string{}
	for barcode := range oriBarcodes {
		barcodes = append(barcodes, barcode)
	}
	for i := 0; i < 16; i++ {
		if i < 8 {
			Primers = append(Primers, Primer{"oriFor_" + strconv.Itoa(i), standardFor + barcodes[i] + oriFor})
		} else {
			Primers = append(Primers, Primer{"oriRev_" + strconv.Itoa(i-8), standardRev + barcodes[i] + oriRev})
		}
	}

	// Build CSV file into stdout
	fmt.Println("Name, Sequence")
	for _, primer := range Primers {
		fmt.Println(primer.Name + "," + primer.Sequence)
	}
}
