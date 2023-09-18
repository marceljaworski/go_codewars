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
	for _, char := range text {
		if strings.Count(counted, string(char)) == 0 {
			counts = append(counts, Tuple{Char: string(char), Count: strings.Count(text, string(char))})
			counted += string(char)
		}
	}
	fmt.Println(len(text), len([]rune(text)))
	return counts
}

func main() {
	fmt.Println(OrderedCount("abracadabra 世 界"))
}
