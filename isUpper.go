package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsUpperCase(s string) bool {
	for _, char := range strings.Join(strings.Fields(s), "") {
		if unicode.IsLower(char) {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(IsUpperCase("WHATDOESTHEFOXsAY"))
	fmt.Println(IsUpperCase("WHAT DOES THE FOX SAY"))
}
