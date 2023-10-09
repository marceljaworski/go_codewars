package main

import (
	"fmt"
	"strings"
)

func ScrabbleScore(st string) int {
	lower := strings.ReplaceAll(strings.ToLower(st), " ", "")
	count := 0
	for _, char := range lower {
		switch string(char) {
		case "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
			count += 1
		case "d", "g":
			count += 2
		case "b", "c", "m", "p":
			count += 3
		case "f", "h", "v", "w", "y":
			count += 4
		case "k":
			count += 5
		case "j", "x":
			count += 8
		case "q", "z":
			count += 10
		}
	}
	return count
}
func main() {
	fmt.Println(ScrabbleScore("cabbage")) // 14
}
