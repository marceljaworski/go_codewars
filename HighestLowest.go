package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	slice := strings.Fields(in)
	max := 0
	min := 9
	for _, char := range slice {
		v, _ := strconv.Atoi(string(char))
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	result := []string{
		strconv.Itoa(max),
		strconv.Itoa(min),
	}
	return strings.Join(result, " ")
}
func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}
