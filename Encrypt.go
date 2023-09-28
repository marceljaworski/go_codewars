package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	if len(text) == 0 {
		return ""
	}
	words := strings.Fields(text)

	for i, word := range words {

		chars := []byte(word)

		if len(chars) > 1 {
			chars[1], chars[len(chars)-1] = chars[len(chars)-1], chars[1]
		}

		words[i] = strconv.Itoa(int(chars[0])) + string(chars[1:])

	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(EncryptThis("Hello"))
	fmt.Println(EncryptThis("good"))
	fmt.Println(EncryptThis("A wise old owl lived in an oak"))
}
