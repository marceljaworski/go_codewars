package main

import (
	"fmt"
)

func HighestRank(nums []int) int {
	myMap := map[int]int{}
	for _, number := range nums {
		myMap[number]++
	}
	max := 0
	solution := 0
	for i, number := range myMap {
		if number >= max {
			max = number
			if solution < i {
				solution = i
			}
			fmt.Println(i, myMap)
		}
		continue
	}
	return solution
}
func main() {
	test := []int{10, 12, 8, 12, 7, 6, 4, 10, 10}
	fmt.Println(HighestRank(test))
}
