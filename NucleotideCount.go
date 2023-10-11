package main

import (
	"errors"
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{'A': 0, 'C': 0, 'T': 0, 'G': 0}
	for _, j := range d {
		// if h[j] != j
		if _, ok := h[j]; !ok {
			return nil, errors.New("invalid nucleotide found")
		}
		h[j]++
	}
	return h, nil
}
func main() {
	var dna DNA = "GATTACA"
	// Counts ist a methode from DNA type
	fmt.Println(dna.Counts())
}
