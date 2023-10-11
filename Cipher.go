package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var cipher bytes.Buffer
	for _, run := range plain {
		rot := run + rune(shiftKey)
		if unicode.IsLetter(run) {
			if unicode.IsLower(run) && rot > 122 || unicode.IsUpper(run) && rot > 90 {
				rot -= 26
			}
			run = rot
		}
		cipher.WriteRune(run)
	}
	return cipher.String()
}
func main() {
	fmt.Println(RotationalCipher("abcdefghijklmnopqrstuvwxyz", 13))
}
