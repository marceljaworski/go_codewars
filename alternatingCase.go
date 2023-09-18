package main

import (
	"fmt"
	"unicode"
)

func ToAlternatingCase(str string) string {
	// var result string = str
	result := make([]rune, 0, len([]rune(str)))
	for _, char := range str {
		fmt.Println(string(char))
		if char == unicode.ToUpper(char) {
			// lower := unicode.ToLower(char)
			result = append(result, unicode.ToLower(char))
		} else {
			upper := unicode.ToUpper(char)
			result = append(result, upper)
		}

	}
	return string(result)
}

func main() {
	fmt.Println(ToAlternatingCase("hAlLo1 örnek iş 世界"))
}
