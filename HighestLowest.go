package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	slice := strings.Fields(in)
	var max, min int
	for i, char := range slice {
		v, _ := strconv.Atoi(string(char))
		if i == 0 {
			max = v
			min = v
		}
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return fmt.Sprintf("%d %d", max, min)
}
func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}
