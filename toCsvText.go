package main

import (
	"fmt"
	"strings"
)

func ToCsvText(array [][]int) string {
	slice := []string{}
	for i := 0; i < len(array); i++ {
		slice2 := []string{}
		for j := 0; j < len(array[i]); j++ {
			slice2 = append(slice2, fmt.Sprint(array[i][j]))
		}
		slice = append(slice, strings.Join(slice2, ","))
	}
	return strings.Join(slice, "\n")
}
func main() {
	text := [][]int{
		{0, 1, 2, 3, 45},
		{10, 11, 12, 13, 14},
		{20, 21, 22, 23, 24},
		{30, 31, 32, 33, 34}}
	fmt.Println(ToCsvText(text))
}
