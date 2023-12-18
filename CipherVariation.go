package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func Cipher(s string, shift int) []string {
	var cipher bytes.Buffer
	for _, run := range s {
		rot := run + rune(shift)
		if unicode.IsLetter(run) {
			if unicode.IsLower(run) && rot > 122 || unicode.IsUpper(run) && rot > 90 {
				rot -= 26
			}
			run = rot
		}
		shift++
		if shift > 26 {
			shift = 1
		}
		cipher.WriteRune(run)
	}
	i := len(cipher.String())
	for i%5 != 0 {
		i++
	}

	slice := make([]string, 5, 5)
	word := ""
	count := 0
	for index, char := range cipher.String() {
		word += string(char)
		if len(word) == i/5 {
			slice[count] = word
			word = ""
			count++
		}
		if index == len(cipher.String())-1 {
			slice[count] = word
		}
	}
	return slice
}
func DeCipher(arr []string, shift int) string {
	var cipher bytes.Buffer
	s := strings.Join(arr, "")
	fmt.Println(s) 
	for _, run := range s {
		rot := run - rune(shift)
		if unicode.IsLetter(run) {
			if unicode.IsLower(run) && rot < 97 || unicode.IsUpper(run) && rot < 65 {
				rot += 26
			}
			run = rot
		}
		if shift > 26 {
			shift = 1
		}
		shift++
		cipher.WriteRune(run)
	}

	return cipher.String()
}

func main() {
	// fmt.Println(RotationalCipher("S mkx bocod", 1))
	fmt.Println(Cipher("I should have known that you would have a perfect answer for me!!!", 2))
	fmt.Println(DeCipher(Cipher("I should have known that you would have a perfect answer for me!!!", 2), 2))

}
