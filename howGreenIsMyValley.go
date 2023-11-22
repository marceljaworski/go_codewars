package main

import (
	"fmt"
	"sort"
)

func MakeValley(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	solution := make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		if len(arr)%2 == 0 {
			if i%2 != 0 {
				solution = append([]int{arr[i]}, solution...)
			} else {
				solution = append(solution, arr[i])
			}
		} else {
			if i%2 == 0 {
				solution = append([]int{arr[i]}, solution...)
			} else {
				solution = append(solution, arr[i])
			}
		}
	}
	return solution
}
func main() {
	fmt.Println(MakeValley([]int{17, 17, 15, 14, 8, 7, 7, 5, 4, 4, 1}))
	fmt.Println(MakeValley([]int{20, 7, 6, 2})) // 20, 6, 2, 7
}
