package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Calc(s string) int {
	a := ""
	for _, char := range s {
		a += strconv.Itoa(int(char))
	}
	b := strings.ReplaceAll(a, "7", "1")

	return Sum(a) - Sum(b)
}
func Sum(s string) (sum int) {

	for _, number := range s {
		sum += int(number)
	}
	return
}

func main() {
	fmt.Println(Calc("abcdef")) // 6
}
