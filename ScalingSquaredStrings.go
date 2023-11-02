package main

import (
	"fmt"
	"strings"
)

func Scale(s string, k, n int) string {
	if s == "" {
		return ""
	}
	slice := strings.Split(s, "\n")
	sol := make([]string, 0, len(slice)*n)
	for _, word := range slice {
		var solutionLetter strings.Builder
		for _, letter := range word {
			for i := 0; i < k; i++ {
				solutionLetter.WriteString(string(letter))
			}
		}
		for j := 0; j < n; j++ {
			sol = append(sol, solutionLetter.String())
		}
	}
	return strings.Join(sol, "\n")
}
func main() {
	fmt.Println(Scale("Kj\nSH", 1, 2))
	fmt.Println(Scale("abcd\nefgh\nijkl\nmnop", 2, 3))
}
