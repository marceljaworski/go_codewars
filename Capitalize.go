package main

import (
	"fmt"
	"strings"
)

func Capitalize(st string) []string {
	var result [2]string
	// result := make([]string, 2)
	for i, char := range st {
		if i%2 == 0 {
			result[0] += strings.ToUpper(string(char))
			result[1] += string(char)
		} else {
			result[1] += strings.ToUpper(string(char))
			result[0] += string(char)
		}
	}
	return result[:]
	// return result
}

func main() {
	fmt.Println(Capitalize("abcdef"))
}
