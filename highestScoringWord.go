package main

import (
	"fmt"
	"strings"
)

type Highest struct {
	word  string
	value int
}

func High(s string) string {
	arr := strings.Fields(s)
	max := Highest{}
	solution := Highest{}
	for _, word := range arr {
		countWord := 0
		for _, char := range word {
			countWord += int(char) - 96
		}
		if countWord > max.value {
			max.value = countWord
			solution.word = word
		}
	}
	return solution.word
}
func main() {
	fmt.Println(High("man i need a taxi up to ubud"))
}
