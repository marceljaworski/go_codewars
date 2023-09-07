package main

import "fmt"

func PositiveSum(nums []int) int {
	arrSum := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			arrSum = arrSum + nums[i]
		}
	}
	return arrSum
}

func main() {
	fmt.Println(PositiveSum([]int{1, 2, 3, 4, -7}))
}
