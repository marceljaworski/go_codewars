package main

import (
	"fmt"
	"strconv"
)

func ExtraPerfect(n int) []int {
	slice := []int{}
	for i := 0; i <= n; i++ {
		bin := strconv.FormatUint(uint64(i), 2)
		if bin[0] == 49 && bin[len(bin)-1] == 49 {
			slice = append(slice, i)
		}
	}
	return slice
}
func main() {
	fmt.Println(ExtraPerfect(7))
}
