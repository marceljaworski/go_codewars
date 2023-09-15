package main

import (
	"fmt"
	"strings"
)

type Tuple struct {
	Char  string
	Count int
}

func OrderedCount(text string) []Tuple {
	counts := make([]Tuple, 0, len([]rune(text)))

	counted := ""
	for _, r := range text {
		if strings.Count(counted, string(r)) == 0 {
			counts = append(counts, Tuple{Char: string(r), Count: strings.Count(text, string(r))})
			counted += string(r)
		}
	}
	return counts
}

func main() {
	fmt.Println(OrderedCount("abracadabra 世 界"))
}
