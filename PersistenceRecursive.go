package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Persistence(n int) int {
	var count int
	str := strings.Split(strconv.Itoa(n), "")
	if len(str) >= 2 {
		_, count = Multiply(str, count)
	}
	return count
}
func Multiply(s []string, count int) (int, int) {
	multiply, _ := strconv.Atoi(s[0])
	for i := 1; i < len(s); i++ {
		num, _ := strconv.Atoi(s[i])
		multiply = multiply * num
	}
	count++
	str := strings.Split(strconv.Itoa(multiply), "")
	if len(str) >= 2 {
		multiply, count = Multiply(str, count)
	}
	return multiply, count
}
func main() {
	fmt.Println(Persistence(39))
}
