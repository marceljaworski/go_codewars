package main

import (
	"fmt"
	"strings"
)

func MultiTable(number int) string {
	slice := make([]string, 0, 10)
	for i := 1; i <= 10; i++ {
		slice = append(slice, fmt.Sprintf("%d * %d = %d", i, number, number*i))
	}
	return strings.Join(slice, "\n")
}
func main() {
	fmt.Println(MultiTable(5))
}
