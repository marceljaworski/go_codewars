package main

import (
	"fmt"
	"strconv"
)

func Second(n int) [2]string {
	var num uint64 = 1 << n
	num = num - 1
	var sol [2]string
	sol[1] = strconv.FormatUint(num, 2)
	result := ""
	if n < 9 {
		sol[0] = strconv.Itoa(n)
	} else {
		for n > 9 {
			result += "9"
			n -= 9
		}
		if n > 0 {
			rest := strconv.Itoa(n)
			result += rest
		}
		sol[0] = result
	}
	return sol
}
func main() {
	fmt.Println(Second(65))
}
