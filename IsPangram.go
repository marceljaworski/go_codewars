package main

import (
	"fmt"
	"regexp"
	"strings"
)

func IsPangram(input string) bool {
	pangram := strings.ReplaceAll(strings.ToLower(input), " ", "")
	re := regexp.MustCompile("[a-z]+").FindAllString(pangram, -1)
	response := strings.Join(re, "")

	myMap := map[string]int{}
	for _, char := range response {
		myMap[string(char)]++
	}
	fmt.Println("test --->", myMap)
	if len(myMap) == 26 {
		return true
	}
	return false
}
func main() {
	fmt.Println(IsPangram("The quick6 brown fox7 jumps8 over the lazy dog"))
}
