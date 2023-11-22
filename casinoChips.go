package main

import (
	"fmt"
	"sort"
)

func Casino(arr []int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	days := 0
	chip := 2
	for chip > 0 {
		chipPerDay := 2
		for i, chips := range arr {
			if chips > 0 {
				arr[i]--
				chipPerDay--
				if arr[i] == 0 {
					chip--
				}
				if chipPerDay == 0 {
					days++
					sort.Slice(arr, func(i, j int) bool {
						return arr[i] > arr[j]
					})
					break
				}
			}

		}
	}
	return days
}
func main() {

	test1 := []int{1, 1, 1}
	test2 := []int{1, 2, 1}
	test3 := []int{8, 2, 8}
	test4 := []int{8, 1, 4}
	test5 := []int{7, 4, 10}
	test6 := []int{12, 12, 12}
	test7 := []int{1, 23, 2}
	fmt.Println(Casino(test1))
	fmt.Println(Casino(test2))
	fmt.Println(Casino(test3))
	fmt.Println(Casino(test4))
	fmt.Println("expected 10. gibt:", Casino(test5)) // 10
	fmt.Println("expected 18. gibt:", Casino(test6))
	fmt.Println(Casino(test7))
}
