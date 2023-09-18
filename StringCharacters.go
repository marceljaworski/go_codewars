package main

import (
	"fmt"
	"unicode"
)

func Solve(s string) []int {
	count := make([]int, 0, 4)
	lower := 0
	upper := 0
	number := 0
	special := 0
	for _, char := range s {

		switch {
		case string(char) != string(unicode.ToUpper(char)):
			lower += 1
		case string(char) != string(unicode.ToLower(char)):
			upper += 1
		case string(char) == "1" || string(char) == "2" || string(char) == "3" || string(char) == "4" || string(char) == "5" || string(char) == "6" || string(char) == "7" || string(char) == "8" || string(char) == "9" || string(char) == "0":
			number += 1
		default:
			special += 1
		}
	}

	count = append(count, upper, lower, number, special)
	return count
	//..
}

func main() {
	fmt.Println(Solve("Codewars@codewars123.com"))
}
