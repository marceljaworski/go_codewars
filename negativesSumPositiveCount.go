package main

import "fmt"

func PositiveCountNegativeSum(nums []int) []int {
	negativeSum := 0
	positiveCount := 0

	for _, num := range nums {
		switch {
		case num <= 0:
			negativeSum = negativeSum + num
		default:
			positiveCount += 1
		}
	}
	return []int{negativeSum, positiveCount}
}

func main() {
	fmt.Println(PositiveCountNegativeSum([]int{1, 2, 3, 4, 8, -7, -8, -10}))
}
