package main

import (
	"fmt"
	"unicode"
	// "strings"
)

func Wave(words string) []string {

	result := make([]string, 0, len(words))
	word := []rune(words)

	fmt.Println(len(word), len(words))
	for i := range word {

		if word[i] == ' ' {
			continue
		}
		word[i] = unicode.ToUpper(word[i])
		result = append(result, string(word))
		word[i] = unicode.ToLower(word[i])
	}

	return result
}

func main() {
	fmt.Println(Wave("hola a"))
}
