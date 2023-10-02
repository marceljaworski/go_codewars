package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FreqSeq(str string, sep string) (r string) {
	count := map[rune]int{}
	slice := []string{}
	for _, char := range str {
		count[char]++
	}
	for _, char := range str {
		slice = append(slice, strconv.Itoa(count[char]))
	}
	return strings.Join(slice, sep)
}

func main() {
	fmt.Println(FreqSeq("hello world", "-"))
}
