package main

import "fmt"

func close_compare(a, b, margin float64) int {

	if a+margin < b {
		return -1
	} else if a > b+margin {
		return 1
	}

	return 0
}

func main() {
	fmt.Println(close_compare(3, 5, 3))
	fmt.Println(close_compare(3, 5, 0))
}
