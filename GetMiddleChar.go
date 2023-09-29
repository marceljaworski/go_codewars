package main

import (
	"fmt"
)

func GetMiddle(s string) string {
	solution := ""
	mid := len(s) / 2
	if len(s)%2 == 0 {
		solution = string(s[mid-1 : mid+1])
	} else {
		solution = string(s[mid])
	}

	return solution
}
func main() {
	fmt.Println(GetMiddle("test"))
	fmt.Println(GetMiddle("testing"))
}
