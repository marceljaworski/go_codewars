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
func MultiplicationTable(size int) (table [][]int) {
	for i := 1; i <= size; i++ {
		arr := make([]int, 0, size)
		for j := 1; j <= size; j++ {
			arr = append(arr, i*j)
		}
		table = append(table, arr)
	}
	return
}
func main() {
	fmt.Println(MultiTable(5))
	fmt.Println(MultiplicationTable(3))
}
