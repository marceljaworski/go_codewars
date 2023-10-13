package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var cipher string
	cipherStr := "zyxwvutsrqponmlkjihgfedcba"
	for _, char := range strings.ToLower(s) {
		if (unicode.IsLetter(char) || unicode.IsDigit(char)) && (len(cipher)+1)%6 == 0 {
			cipher += " "
		}
		if unicode.IsLetter(char) {
			rot := int(char) - 97
			cipher += string(cipherStr[rot])
		}
		if unicode.IsDigit(char) {
			cipher += string(char)
		}
	}
	return cipher
}

func main() {
	fmt.Println(Atbash("test"))
}
