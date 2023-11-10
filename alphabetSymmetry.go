package main

import (
	"fmt"
	"strings"
)

func Alphbet(slice []string) []int {
	solution := make([]int, 0, len(slice))
	for _, word := range slice {
		count := 0
		for i, char := range strings.ToLower(word) {
			if int(char)-97 == i {
				count++
			}
		}
		solution = append(solution, count)
	}
	return solution
}
func main() {
	s := []string{"abode", "ABc", "xyzD"}
	fmt.Println(Alphbet(s))
}
