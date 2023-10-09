package main

import (
	"fmt"
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	myMap := map[string]int{}
	for i, value := range in {
		for j := 0; j < len(value); j++ {
			char := strings.ToLower(value[j])
			myMap[char] += i
		}
	}
	return myMap
}
func main() {
	input := map[int][]string{
		1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  {"D", "G"},
		3:  {"B", "C", "M", "P"},
		4:  {"F", "H", "V", "W", "Y"},
		5:  {"K"},
		8:  {"J", "X"},
		10: {"Q", "Z"}}
	fmt.Println(Transform(input))
}
