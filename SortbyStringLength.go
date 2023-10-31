package main

import (
	"fmt"
)

func SortByLength(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			fmt.Println(i, j)
			if len(arr[j]) < len(arr[i]) {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
func main() {
	fmt.Println(SortByLength([]string{"Telescopes", "Glasses", "Eyes", "Monocles"}))
}
