package main

import (
	"fmt"
	"strings"
	"unicode"
)

func solve(str string) string {
	// Your code here and happy coding!
	lower, upper := 0, 0

	for _, char := range str {
		if string(char) != string(unicode.ToUpper(char)) {
			lower++
		} else {
			upper++
		}
	}
	if upper == lower || upper < lower {
		return strings.ToLower(str)
	} else {
		return strings.ToUpper(str)
	}

}
func main() {
	fmt.Println(solve("coDe"))
	fmt.Println(solve("CODE"))
	fmt.Println(solve("coDE"))
}
