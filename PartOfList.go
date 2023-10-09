package main

import (
	"fmt"
	"strings"
)

func PartList(arr []string) string {
	solution := make([]string, 0, len(arr))
	for i := 0; i < len(arr)-1; i++ {
		solution = append(solution, "("+strings.Join(arr[0:i+1], " ")+", "+strings.Join(arr[i+1:], " ")+")")
	}
	return strings.Join(solution, "")
}

func main() {
	test := []string{"az", "toto", "picaro", "zone", "kiwi"}
	fmt.Println(PartList(test))
}
