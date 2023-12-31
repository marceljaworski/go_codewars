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

func DecodeRotationalCipher(plain string, shiftKey int) string {
	var cipher bytes.Buffer
	for _, run := range plain {
		rot := run - rune(shiftKey)
		if unicode.IsLetter(run) {
			if unicode.IsLower(run) && rot < 97 || unicode.IsUpper(run) && rot < 65 {
				rot += 26
			}
			run = rot
		}
		cipher.WriteRune(run)
	}
	return cipher.String()
}
func main() {
	fmt.Println(RotationalCipher("abcdefghijklmnopqrstuvwxyz", 13))
	fmt.Println(RotationalCipher("iamapandabear", 3))
	fmt.Println(DecodeRotationalCipher("ldpdsdqgdehdu", 3))

}
