package main

import (
	"fmt"
	"strings"
)

func MobileKeyboard(str string) int {
	result := 0
	for _, char := range str {
		switch {
		case strings.Contains("adgjmptw", string(char)):
			result = result + 2
		case strings.Contains("behknqux", string(char)):
			result = result + 3
		case strings.Contains("cfilorvy", string(char)):
			result = result + 4
		case strings.Contains("sz", string(char)):
			result = result + 5
		default:
			result++
		}
	}
	return result
}
func main() {
	fmt.Println(MobileKeyboard("codewars"))
}
