package main

import "fmt"

func InvertValues(arr []int) []int {
	invArr := make([]int, len(arr))
	for i, num := range arr {
		invArr[i] = -num
	}
	return invArr
}

func main() {
	fmt.Println(InvertValues([]int{1, 2, 3, 4, 8, -7, -8, -10}))
}
