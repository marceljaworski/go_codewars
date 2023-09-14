// Simple, given a string of words, return the length of the shortest word(s).
package main

import (
	"fmt"
	"strings"
)

func FindShort2(s string) int {
	words := strings.Fields(s)

	result := -1

	for _, word := range words {
		l := len(word)
		if result < 0 || l < result {
			result = l
		}
	}

	return result
}

func FindShort(s string) int {

	words := strings.Split(s, " ")

	var values []int

	for _, word := range words {
		values = append(values, len(word))

	}

	min := values[0]
	for _, number := range values {
		if number < min {
			min = number
		}
	}
	fmt.Println(words, values)
	return min
}
func main() {
	fmt.Println(FindShort("hallo  du und alles suzammen"))
}
