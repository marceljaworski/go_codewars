package main

import (
	"fmt"
	"sort"
	"strings"
)

func Solve(s string) bool {
	x := strings.Split(s, "")
	sort.Strings(x)
	str := strings.Join(x, "")
	for i := 1; i < len(str); i++ {
		fmt.Println(str[i-1]+1, str[i])
		if str[i-1]+1 != str[i] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(Solve("dabc"))
}
