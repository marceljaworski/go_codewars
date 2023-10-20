package main

import (
	"fmt"
	"strings"
)

func WordPattern(word string) string {
	var count int
	var s strings.Builder
	d := ""
	neuWord := strings.ToLower(word)
	myMap := map[rune]int{}
	for _, char := range neuWord {
		if _, ok := myMap[char]; !ok {
			myMap[char] = count
			count++
		}
		fmt.Fprintf(&s, "%s%d", d, myMap[char])
		d = "."
	}
	fmt.Println(myMap)
	return s.String()
}
func main() {
	fmt.Println(WordPattern("Hippopotomonstrosesquippedaliophobia"))
}
