package main

import "fmt"

func DNAStrand(dna string) string {
	complement := ""
	for _, symbol := range dna {
		switch string(symbol) {
		case "A":
			complement += "T"
		case "T":
			complement += "A"
		case "C":
			complement += "G"
		case "G":
			complement += "C"
		}
	}

	return complement
}

func main() {
	fmt.Println(DNAStrand("ATTGC"))
	fmt.Println(DNAStrand("GTAT"))
}
